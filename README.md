### Step to compile and run webapp on local machine
1. cmd -> rm -rf ./node_modules && npm install
2. npm run build:sass
3. go run .

### Command often uses
1. cf apps
2. cf routes
3. cf api -> cf api  https://api.sys.adp.ec1.aws.aztec.cloud.allianz --skip-ssl-validation
4. cf scale -> ex) cf scale app_name -i 1
5. cf push -> ex) cf push app_name -p ./go-website -m 256M -n sub_domain-temp --no-start, cf push app_name -m 256M --random-route --no-start
6. cf logs app_name --recent
7. curl https://go-webapp.apps.adp.ec1.aws.aztec.cloud.allianz

### Step to deploy webapp over pcf
1. cf login and choose your space
2. cf push app_name -b https://github.com/cloudfoundry/go-buildpack.git -u http -m 128M --no-start, 
cf push go-webapp -b https://github.com/cloudfoundry/go-buildpack.git -u http --no-manifest
3. cf push go-webapp-v2 -n go-webapp-temp -b https://github.com/cloudfoundry/go-buildpack.git -u http --no-manifest -m 256M --no-start
4. cf map-route go-webapp-v2 apps.adp.ec1.aws.aztec.cloud.allianz -n go-webapp
5. cf unmap-route go-webapp apps.adp.ec1.aws.aztec.cloud.allianz -n go-webapp
6. cf unmap-route go-webapp-2 apps.adp.ec1.aws.aztec.cloud.allianz -n go-webapp-temp
   
### Scratch version
1. 
2. 

### Book references
`
https://morrisjs.github.io/morris.js/index.html#getting-started
https://material.io/develop/web/components/layout-grid/
https://docs.run.pivotal.io/devguide/deploy-apps/environment-variable.html
https://cli.cloudfoundry.org/en-US/cf/push.html
https://docs.cloudfoundry.org/buildpacks/go/index.html
`