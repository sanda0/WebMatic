
import { BadgePlus, Import } from "lucide-react";
import Button from "../../components/Button";
import { Outlet } from "react-router-dom";

export default function MainLayout() {
	return (
		<div className="h-screen w-full bg-slate-300 flex">
			<div className="h-screen w-[300px] bg-slate-700 text-white">
        <ul className="p-4 ">
          <li ><div className="text-center text-4xl">WebMatic</div></li>
          <li className="mt-4">
            <Button> <BadgePlus></BadgePlus> New Matic </Button>
          </li>
          <li className="mt-2">
            <Button> <Import></Import> Import Matic </Button>
          </li>
        </ul>
      </div>
      <div className="h-screen w-full">
        <Outlet></Outlet>
      </div>
		</div>
	);
}
