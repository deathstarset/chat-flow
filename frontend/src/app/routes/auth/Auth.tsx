import { Outlet } from "react-router-dom";

export default function Auth() {
  return (
    <div className="container mx-auto h-[100vh]">
      <div className="h-[8vh] flex items-center">
        <h1 className="text-3xl font-bold">Chatflow</h1>
      </div>
      <div className="items-center flex justify-center h-[80vh]">
        <Outlet />
      </div>
    </div>
  );
}
