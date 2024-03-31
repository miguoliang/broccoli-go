import { Field, Form, Formik } from "formik";
import * as Yup from "yup";
import { DtoCreateVertexRequest, useCreateVertex } from "./api.ts";

const CreateVertexSchema = Yup.object().shape({
  name: Yup.string().required("Name is required"),
  type: Yup.string().required("Type is required"),
});

const NewVertexForm = () => {
  const mutation = useCreateVertex();

  return (
    <Formik<DtoCreateVertexRequest>
      initialValues={{
        name: "",
        type: "",
      }}
      validationSchema={CreateVertexSchema}
      validateOnBlur={false}
      validateOnChange
      onSubmit={(values) => mutation.mutate({ data: values })}
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
};

export default NewVertexForm;
