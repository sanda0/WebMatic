
import { BadgePlus, Import } from "lucide-react";
import Button from "../../components/Button";
import { Outlet } from "react-router-dom";

export default function MainLayout() {
	return (
		<div className="flex w-full h-screen bg-slate-300">
			<div className="h-screen w-[300px] bg-slate-700 text-white">
        <ul className="p-4 ">
          <li ><div className="text-4xl text-center">WebMatic</div></li>
          <li className="mt-4">
            <Button onClick={(e)=>{alert("hello")}}> <BadgePlus></BadgePlus> New Matic </Button>
          </li>
          <li className="mt-2">
            <Button> <Import></Import> Import Matic </Button>
          </li>
        </ul>
      </div>
      <div className="w-full h-screen">
        <Outlet></Outlet>
      </div>
		</div>
	);
}
