
import { BadgePlus, Home, Import } from "lucide-react";
import Button from "../../components/Button";
import { Outlet, useNavigate } from "react-router-dom";

export default function MainLayout() {
  const navigate = useNavigate()
  


	return (
		<div className="flex w-full h-screen p-0 m-0 bg-slate-300">
			<div className="h-screen w-[300px] bg-slate-700 text-white">
        <ul className="p-4 ">
          <li ><div className="text-4xl text-center">WebMatic</div></li>
          <li className="mt-4">
            <Button onClick={(e)=>{navigate("/")}}> <Home></Home> Home </Button>
          </li>
          <li className="mt-2">
            <Button onClick={(e)=>{navigate("/new")}}> <BadgePlus></BadgePlus> New Matic </Button>
          </li>
          <li className="mt-2">
            <Button> <Import></Import> Import Matic </Button>
          </li>
        </ul>
      </div>
      <div className="w-full h-screen p-0 m-0" style={{paddingTop:10}}>
        <Outlet></Outlet>
      </div>
		</div>
	);
}
