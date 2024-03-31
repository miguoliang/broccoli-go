module.exports = {
  broccoli: {
    input: "../docs/v3/openapi/openapi.yaml",
    output: {
      target: "src/api.ts",
      baseUrl: "/api",
      client: "react-query",
      override: {
        mutator: {
          path: "src/axios.ts",
          name: "customInstance",
        },
      },
    },
  },
};