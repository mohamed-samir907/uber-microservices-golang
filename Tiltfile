load('ext://restart_process', 'docker_build_with_restart')

# ---------------------------
#       <k8s config>
# ---------------------------

k8s_yaml('./infra/development/k8s/app-config.yaml')


# ---------------------------
#       <API Gateway>
# ---------------------------

gateway_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./services/api-gateway'

local_resource(
    'api-gateway-compile',
    gateway_compile_cmd,
    deps=['./services/api-gateway', './shared'],
    labels='compiles')

docker_build_with_restart(
    'ride-sharing/api-gateway',
    context='.',
    entrypoint=['/app/build/api-gateway'],
    dockerfile='./infra/development/docker/api-gateway.Dockerfile',
    only=[
        './build/api-gateway',
        './shared',
    ],
    live_update=[
        sync('./build', '/app/build'),
        sync('./shared', '/app/shared'),
    ],
)

k8s_yaml('./infra/development/k8s/api-gateway-deployment.yaml')
k8s_resource('api-gateway',
    port_forwards=8081,
    resource_deps=['api-gateway-compile'],
    labels="services",
)

# ---------------------------
#       <Trip Service>
# ---------------------------

trip_service_compile_cmd = 'GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o build/trip-service ./services/trip-service/cmd/main.go'

local_resource(
    'trip-service-compile',
    trip_service_compile_cmd,
    deps=['./services/trip-service', './shared'],
    labels="compiles",
)

docker_build_with_restart(
    'ride-sharing/trip-service',
    context='.',
    dockerfile='./infra/development/docker/trip-service.Dockerfile',
    entrypoint=['/app/build/trip-service'],
    only=[
        './build/trip-service',
        './shared'
    ],
    live_update=[
        sync('./build', '/app/build'),
        sync('./shared', '/app/shared'),
    ]
)

k8s_yaml('./infra/development/k8s/trip-service-deployment.yaml')
k8s_resource(
    'trip-service',
    port_forwards=8083,
    resource_deps=['api-gateway-compile'],
    labels='services',
)

# ---------------------------
#       <Web Frontend>
# ---------------------------

docker_build(
    'ride-sharing/web',
    context='.',
    dockerfile="./infra/development/docker/web.Dockerfile",
)

k8s_yaml('./infra/development/k8s/web-deployment.yaml')
k8s_resource(
    'web',
    port_forwards=3000,
    labels="frontend",
)
