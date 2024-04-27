module.exports = {
  broccoli: {
    input: "../api/openapi.yaml",
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