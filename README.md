# otel-grpc-example
gRPCサーバーでの分散トレーシング等のデモ<br>
このブランチは、opentelemetryを導入する前のuidサービスとtaskサービスとzipkinが起動するだけの状態のブランチです．<br>

## 起動方法
docker desktopでkubernetesを起動していることが前提になっています．<br>
1. skaffoldコマンドのインストール<br>
   ```
   brew install skaffold
   ```
2. 実行<br>
   ```
   cd src/services
   skaffold dev
   ```
3. ブラウザでzipkinにアクセス
    ```
    http://localhost:59411/zipkin/
    ```
4. uidサービスにリクエストを送る
   ```
   grpcurl -plaintext localhost:55051 moa.api.uid.v1.UIDService.Generate
   ```
   想定されるレスポンスは以下の通りです．
   ```
   {
     "ID": "7a75ecd2-90cd-4f7c-ae8c-b7d4f4f7e7b6"
   }
   ```
5. taskサービスにリクエストを送る
   ```
   grpcurl -plaintext -d '{"task":{"text":"test"}}' localhost:55050 moa.api.task.v1.TaskService.Create
   ```
   想定されるレスポンスは以下の通りです．(taskの時は小文字のid)
   ```
   {
     "id": "39d6533c-5317-4c5b-b60f-6e47070240a9"
   }
   ```