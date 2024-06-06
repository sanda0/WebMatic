export default function Button(props) {
	return (
		<button onClick={(e)=>{props.onClick(e)}} className="flex w-full gap-2 p-2 m-1 text-center text-white border rounded-md hover:bg-white hover:text-slate-700 ">
			{props.children}
		</button>
	);
}
