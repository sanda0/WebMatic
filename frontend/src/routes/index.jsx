import {
  Route,
	createBrowserRouter,
	createRoutesFromElements,
} from "react-router-dom";
import MainLayout from "../pages/layouts/MainLayout";
import NewMatic from "../pages/NewMatic";
import Editor from "../pages/Editor";
import Home from "../pages/Home";

const router = createBrowserRouter(createRoutesFromElements(
  <Route path="/" element={<MainLayout></MainLayout>}>
    <Route path="" element={<Home></Home>}></Route>
    <Route path="new" element={<NewMatic></NewMatic>}></Route>
    <Route path="matic/:id" element={<Editor></Editor>}></Route>
  </Route>
));

export default router;
