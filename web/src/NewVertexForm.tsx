import { DefaultError, useMutation } from "@tanstack/react-query";
import { CreateVertexRequest, CreateVertexResponse } from "./dto.tsx";
import axios from "axios";
import { Field, Form, Formik } from "formik";
import * as Yup from "yup";

const CreateVertexSchema = Yup.object().shape({
  name: Yup.string().required("Name is required"),
  type: Yup.string().required("Type is required"),
});

export default function NewVertexForm() {
  const mutation = useMutation<
    CreateVertexResponse,
    DefaultError,
    CreateVertexRequest
  >({
    mutationKey: ["createVertex"],
    mutationFn: async (variables) => {
      const response = await axios.post<CreateVertexResponse>(
        "/api/vertex",
        variables,
      );
      return response.data;
    },
  });

  return (
    <Formik<CreateVertexRequest>
      initialValues={{
        name: "",
        type: "",
      }}
      validationSchema={CreateVertexSchema}
      validateOnBlur={false}
      validateOnChange
      onSubmit={(values) => mutation.mutate(values)}
    >
      {({ errors, touched }) => (
        <Form>
          <Field name="name" type="text" placeholder="Name" />
          {errors.name && touched.name ? <div>{errors.name}</div> : null}
          <Field name="type" type="text" placeholder="Type" />
          {errors.type && touched.type ? <div>{errors.type}</div> : null}
          <button type="submit" disabled={mutation.isPending}>
            Create
          </button>
        </Form>
      )}
    </Formik>
  );
}
