Monolithic:
src/
- main.go         // fetch configs and build an app by calling build() and run it
- app.go          // composed of services. build dependencies in build() method
- config.go       // configuration managment from a yaml file or env variables.
- migrations/
- models/
  - foo1.go
  - foo2.go
- repository/
  - foo1Repo.go   // interface
  - foo2Repo.go 
  - foo1RepoStorage1.go // concrete
-  services/
  - service.go    // interface
  - foo1Service.go

Microservices: Above pattern for each domain in separate folder.

Design Choice: A repository for each model or a common repository for entire domain?

    
