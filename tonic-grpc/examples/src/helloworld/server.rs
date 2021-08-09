#[macro_use]
extern crate lazy_static;
extern crate rbatis;

use log::info;
use rbatis::crud::CRUD;
use rbatis::crud_table;
use rbatis::plugin::page::{Page, PageRequest};
use rbatis::rbatis::Rbatis;

use hello_world::{HelloReply, HelloRequest};
use hello_world::greeter_server::{Greeter, GreeterServer};
use tonic::{Request, Response, Status, transport::Server};

// init global rbatis pool
lazy_static! {
    static ref RB: Rbatis = Rbatis::new();
}

#[crud_table]
#[derive(Clone, Debug)]
pub struct MemberInfo {
    pub id: Option<i64>,
    pub name: Option<String>,
    pub avatar: Option<String>,
}

pub mod hello_world {
    tonic::include_proto!("helloworld");
}

#[derive(Default)]
pub struct MyGreeter {}

#[tonic::async_trait]
impl Greeter for MyGreeter {
    async fn say_hello(
        &self,
        request: Request<HelloRequest>,
    ) -> Result<Response<HelloReply>, Status> {
        println!("Got a request from {:?}", request.remote_addr());
        let name: String = String::from(&request.get_ref().name);
        println!("name {:?}", name);

        let req = PageRequest::new(1, 20); //分页请求，页码，条数
        let wrapper = RB
            .new_wrapper()
            .is_not_null("avatar")
            .ne("name", name.as_str())
            .order_by(true, &vec!["id"]);
        let data: Page<MemberInfo> = RB.fetch_page_by_wrapper(&wrapper, &req).await.unwrap();

        let reply = hello_world::HelloReply {
            message: serde_json::json!(data).to_string(),
        };
        Ok(Response::new(reply))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    //  RUSTFLAGS="-C target-cpu=native" cargo build --release
    let adder = "[::1]:50051".parse().unwrap();
    let greeter = MyGreeter::default();

    //启用日志输出
    fast_log::init_log("requests.log", 1000, log::Level::Info, None, true).unwrap();
    info!("service [::1]:50051 start");
    //初始化连接池
    RB.link("mysql://root:@127.0.0.1:3306/pirate_tv")
        .await
        .unwrap();

    println!("GreeterServer listening on {}", adder);

    Server::builder()
        .add_service(GreeterServer::new(greeter))
        .serve(adder)
        .await?;

    Ok(())
}
