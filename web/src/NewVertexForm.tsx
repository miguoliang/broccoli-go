import { Field, Form, Formik } from "formik";
import { object, string } from "yup";
import { useCreateVertex } from "./api.ts";

const CreateVertexSchema = object({
  industry: string().required("Industry is required"),
  catalog: string().required("Catalog is required"),
  name: string().required("Name is required"),
});

const NewVertexForm = () => {
  const mutation = useCreateVertex();

  return (
    <div className={"max-w-7xl mx-auto p-2"}>
      <Formik
        initialValues={{
          industry: "",
          name: "",
          catalog: "",
        }}
        validationSchema={CreateVertexSchema}
        validateOnBlur={false}
        validateOnChange
        onSubmit={(values) => mutation.mutate({
          data: {
            name: values.name,
            type: "product",
            properties: {
              industry: values.industry,
              catalog: values.catalog,
            },
          },
        })}
      >
        {({ errors, touched }) => (
          <Form className={"grid grid-cols-10 gap-2 items-center"}>
            <label className={"font-bold text-right"} htmlFor="industry">Industry</label>
            <Field className={"rounded-md border p-2 col-span-2"} id="industry" name="industry" type="text"
                   placeholder="Industry" />
            <label className={"font-bold text-right"} htmlFor="productName">Product Name</label>
            <Field className={"rounded-md border p-2 col-span-2"} id="productName" name="productName" type="text"
                   placeholder="Product Name" />
            {errors.name && touched.name ? <div>{errors.name}</div> : null}
            <label className={"font-bold text-right"} htmlFor="productType">Product Type</label>
            <Field className={"rounded-md border p-2 col-span-2"} id={"productType"} name="productType" type="text"
                   placeholder="Type" />
            {errors.catalog && touched.catalog ? <div>{errors.catalog}</div> : null}
            <div className={"text-right"}>
              <button className={"border p-2 rounded-md"} type="submit" disabled={mutation.isPending}>
                Create
              </button>
            </div>
          </Form>
        )}
      </Formik>
    </div>
  );
};

export default NewVertexForm;
