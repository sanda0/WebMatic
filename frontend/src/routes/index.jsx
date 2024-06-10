import {
  Route,
	createBrowserRouter,
	createRoutesFromElements,
} from "react-router-dom";
import MainLayout from "../pages/layouts/MainLayout";
import NewMatic from "../pages/NewMatic";
import Editor from "../pages/Editor";

const router = createBrowserRouter(createRoutesFromElements(
  <Route path="/" element={<MainLayout></MainLayout>}>
    <Route path="new" element={<NewMatic></NewMatic>}></Route>
    <Route path="matic/:name" element={<Editor></Editor>}></Route>
  </Route>
));

export default router;
