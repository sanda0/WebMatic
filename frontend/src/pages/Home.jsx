import { useEffect, useState } from "react";
import { GetAllMatics } from "../../wailsjs/go/main/App";
import MaticItem from "../components/MaticItem";

export default function Home() {
  const [matics,setMatic] = useState([])

  useEffect(()=>{
    GetAllMatics().then((res)=>{
      console.log(res)
      if(res.status == 200) {
        setMatic(res.data)
      }
    })
    
  },[])

  function onEdit(file){
    console.log(file)
  }

	return (
		<div className="content-center h-screen">
			<div className=" rounded-md w-[60%] p-5 h-[80%]   m-auto text-slate-700 overflow-y-auto">
				<div className="text-3xl text-center"> Matics</div>
				<div className="grid grid-cols-1 gap-4 mt-4">
            {matics.map((e,i)=>{
              console.log(e);
              return <MaticItem key={i} name={e.name} id={e.ID} onEdit={(f)=>onEdit}></MaticItem>
            })}
        </div>
			</div>
		</div>
	);
}
