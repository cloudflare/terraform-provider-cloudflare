# Debugging with Delve

In the past, the only way to know what was happening in Terraform was to drop 
something like this around your code.

```
log.Println("[DEBUG] Something happened!")log.Printf("[DEBUG] broken thing: %#v", thing)
```

This is incredibly tedious, and you end up wasting time going back and forth on 
what you want to inspect and in what states.

Since we've finally upgraded to `terraform-plugin-sdk` v2, we're able to advance
from log based debugging with Terraform to using a debugger. My choice for 
Golang is [Delve] however some Gophers will favour [gdb] instead.

Below, I'm using VS Code as my editor however should be able to sub in your 
editor equivalent and follow along.

- Drop into a terminal session
- Clone the git repository into your Go directory
  ```
  $ git clone git@github.com:cloudflare/terraform-provider-cloudflare.git /Users/jacob/go/src/github.com/cloudflare/
  ```
- Build the development version of the provider binary. It is important to use 
  the development version to disable compiler optimisations and inlining; without 
  these compile flags some parts of Delve just won't work.
  ```
  $ make build-dev
  ```
- Once the build finishes, confirm the binary exists
  ```
  $ file terraform-provider-cloudflare_99.0.0
    terraform-provider-cloudflare_99.0.0: Mach-O 64-bit executable x86_64
  ```
- Now that the binary is created, you can execute Delve in headless mode. I use 
  a static listen address here to make the debugging configuration consistent but 
  you can also use the ephemeral port assigned if you'd rather do that.
  ```
  $ dlv exec --headless ./terraform-provider-cloudflare_99.0.0 --listen 127.0.0.1:44444 -- --debug
  API server listening at: 127.0.0.1:44444
  debugserver-@(#)PROGRAM:LLDB  PROJECT:lldb-1300.0.32.2
    for x86_64.
  Got a connection, launched process ./terraform-provider-cloudflare_99.0.0 (pid = 92684).
  ```
  Be sure to leave this window/tab open. It will soon have an environment 
  variable configuration for you to prefix your Terraform commands with.
- In VS Code, create a new `launch.json` for debugging. Drop in the following snippet.
  ```json
    {
      "version": "0.2.0",
      "configurations": [
        {
          "name": "Terraform delve debugger",
          "type": "go",
          "request": "attach",
          "mode": "remote",
          "remotePath": "${workspaceFolder}",
          "port": 44444,
          "host": "127.0.0.1",
          "apiVersion": 1.0,
        }
      ]
    }
    ```
- You can now hit the "start debugger" button in VS code.

  ![GdrJ0Q5l-CYg5yyH6-jwJ1Oeee](https://user-images.githubusercontent.com/283234/134456154-55f06831-5016-46b1-ae7f-563baf5ddc36.png)
  
- If you've done everything correct up until this point, you should now be able 
  to pop back to the Delve terminal session and see it has now output a 
  `TF_REATTACH_PROVIDERS` configuration below your `dlv exec` command.
  ```
  {"@level":"debug","@message":"plugin address","@timestamp":"2021-09-23T13:56:14.703121+10:00","address":"/var/folders/36/zlscnhfn27n1yxx52cr1kdmr0000gp/T/plugin1090848880","network":"unix"}
  Provider started, to attach Terraform set the TF_REATTACH_PROVIDERS env var:

    TF_REATTACH_PROVIDERS='{"registry.terraform.io/cloudflare/cloudflare":{"Protocol":"grpc","ProtocolVersion":5,"Pid":92684,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/36/zlscnhfn27n1yxx52cr1kdmr0000gp/T/plugin1090848880"}}}'

  {"@level":"debug","@message":"stdio EOF, exiting copy loop","@timestamp":"2021-09-23T13:58:24.698711+10:00"}
  {"@level":"debug","@message":"stdio EOF, exiting copy loop","@timestamp":"2021-09-23T13:58:24.698712+10:00"}
  ```
- Now, it is just a matter of prefixing your Terraform commands (apply/plan, 
  go tests) with this environment variable and it will hit the breakpoints 
  you've set. You can also export this into the environment if you'd rather 
  instead of prefixing all commands.
  ```
  $ TF_REATTACH_PROVIDERS='{"registry.terraform.io/cloudflare/cloudflare":{"Protocol":"grpc","ProtocolVersion":5,"Pid":92684,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/36/zlscnhfn27n1yxx52cr1kdmr0000gp/T/plugin1090848880"}}}' \  
    CLOUDFLARE_EMAIL="..." CLOUDFLARE_API_KEY="..." terraform apply
  ```

[Delve]: https://github.com/go-delve/delve
[gdb]: https://golang.org/doc/gdb
