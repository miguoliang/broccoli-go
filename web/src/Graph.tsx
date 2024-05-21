import {
  Button,
  Description,
  Dialog,
  DialogPanel,
  DialogTitle,
  Transition,
  TransitionChild,
} from "@headlessui/react";
import { ErrorMessage, Field, FieldArray, Form, Formik } from "formik";
import { useState } from "react";
import { MdDelete } from "react-icons/md";
import ReactFlow, {
  Background,
  BackgroundVariant,
  Controls,
  MiniMap,
  Panel,
  ReactFlowProvider,
} from "reactflow";
import * as Yup from "yup";
import SearchBox from "./components/SearchBox";
import useGraphStore from "./stores/GraphStore";

const Graph = () => {
  const { nodes, edges } = useGraphStore();
  const [isOpen, setIsOpen] = useState(false);
  return (
    <>
      <ReactFlowProvider>
        <ReactFlow nodes={nodes} edges={edges}>
          <Panel position="top-left">
            <div className="flex justify-between text-sm shadow-around bg-white py-1">
              <SearchBox />
              <button className="px-2" onClick={() => setIsOpen(true)}>
                Add Node
              </button>
            </div>
          </Panel>
          <Controls className="bg-white" />
          <MiniMap />
          <Background variant={BackgroundVariant.Dots} />
        </ReactFlow>
      </ReactFlowProvider>
      <AddNodeDialog isOpen={isOpen} onClose={() => setIsOpen(false)} />
    </>
  );
};

type AddNodeDialogProps = {
  isOpen: boolean;
  onClose: () => void;
};

const AddNodeDialog = ({ isOpen, onClose }: Readonly<AddNodeDialogProps>) => {
  return (
    <Transition appear show={isOpen} as="div">
      <Dialog
        as="div"
        className="relative z-10 focus:outline-none"
        onClose={onClose}
      >
        {/* The backdrop, rendered as a fixed sibling to the panel container */}
        <TransitionChild
          as="div"
          className="fixed inset-0 bg-black/10"
          aria-hidden="true"
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        />
        {/* Full-screen container to center the panel */}
        <div className="fixed inset-0 flex w-screen items-center justify-center p-4">
          {/* The actual dialog panel  */}

          <TransitionChild
            as={DialogPanel}
            className="max-w-xl space-y-4 bg-white p-12 rounded-md shadow-around"
            enter="ease-out duration-300"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <DialogTitle className="font-bold">Deactivate account</DialogTitle>
            <Description>
              This will permanently deactivate your account
            </Description>
            <AddNodeForm />
          </TransitionChild>
        </div>
      </Dialog>
    </Transition>
  );
};

interface AddNodeFormValues {
  name: string;
  type: string;
  props: Array<{ key: string; value: string }>;
}

const validationSchema = Yup.object({
  name: Yup.string().required("Name is required"),
  type: Yup.string().required("Type is required"),
  props: Yup.object().shape({
    prop1: Yup.string().required("Prop 1 is required"),
    prop2: Yup.string().required("Prop 2 is required"),
    // Add more prop validations as needed
  }),
});

const initialValues: AddNodeFormValues = {
  name: "",
  type: "",
  props: [],
};

const AddNodeForm = () => {
  const handleSubmit = (values: AddNodeFormValues) => {
    // Handle form submission
    console.log(values);
  };

  return (
    <Formik
      initialValues={initialValues}
      validationSchema={validationSchema}
      onSubmit={handleSubmit}
    >
      {({ values: { props } }) => (
        <Form>
          <div>
            <label htmlFor="name">Name</label>
            <Field type="text" id="name" name="name" />
            <ErrorMessage name="name" component="div" />
          </div>
          <div>
            <label htmlFor="type">Type</label>
            <Field type="text" id="type" name="type" />
            <ErrorMessage name="type" component="div" />
          </div>
          <FieldArray name="props">
            {({ push, remove }) => (
              <div>
                <h3>Props</h3>
                {props.map((_, index) => (
                  <div
                    key={index}
                    className="flex gap-1 hover:bg-gray-100 p-1 rounded-sm"
                  >
                    <Field
                      type="text"
                      id={`props.${index}.key`}
                      name={`props.${index}.key`}
                      placeholder="Key"
                      className="flex-grow-1 appearance-none border-none rounded-sm focus:ring-0 outline-none"
                    />
                    <ErrorMessage name={`props.${index}.key`} component="div" />
                    <Field
                      type="text"
                      id={`props.${index}.value`}
                      name={`props.${index}.value`}
                      placeholder="Value"
                      className="flex-grow-1 appearance-none border-none rounded-sm focus:ring-0 outline-none"
                    />
                    <ErrorMessage
                      name={`props.${index}.value`}
                      component="div"
                    />
                    <Button className="flex-grow-0" type="button" onClick={() => remove(index)}>
                      <MdDelete />
                    </Button>
                  </div>
                ))}
                <Button
                  type="button"
                  className="px-1"
                  onClick={() => push({ key: "", value: "" })}
                >
                  Add Prop
                </Button>
              </div>
            )}
          </FieldArray>
          {/* Add more prop fields as needed */}
          <Button
            className="inline-flex items-center gap-2 rounded-md bg-gray-700 py-1.5 px-3 text-sm/6 font-semibold text-white shadow-inner shadow-white/10 focus:outline-none data-[hover]:bg-gray-600 data-[open]:bg-gray-700 data-[focus]:outline-1 data-[focus]:outline-white"
            onClick={close}
            type="submit"
          >
            Got it, thanks!
          </Button>
        </Form>
      )}
    </Formik>
  );
};

export default Graph;
