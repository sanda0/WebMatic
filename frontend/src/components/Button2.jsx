export default function Button2(props) {
	return (
		<button onClick={(e)=>{props.onClick(e)}} className="flex items-center w-full gap-2 p-2 m-1 text-center border-2 rounded-md text-slate-700 border-slate-700 hover:bg-white hover:text-slate-700 ">
			{props.children}
		</button>
	);
}
