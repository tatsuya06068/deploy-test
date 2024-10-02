package main

import (
    "context"
    "log"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    pb "path/to/your/proto" // プロトコルバッファファイルから生成されたパッケージをインポート
)

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()

    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
    if err != nil {
        log.Fatalf("Failed to start HTTP server: %v", err)
    }

    log.Println("gRPC-Gateway server listening on :8080")
    http.ListenAndServe(":8080", mux)
}
