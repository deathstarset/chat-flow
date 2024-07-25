import { Link } from "react-router-dom";
import LoginForm from "./LoginForm";

export default function Login() {
  return (
    <div className="card bg-base-100 w-96 shadow-xl">
      <div className="card-body flex flex-col gap-4">
        <h2 className="card-title">Login</h2>
        <LoginForm />
        <span className="text-center text-sm">
          Don't have an account ?{" "}
          <Link
            to="/auth/register"
            className="text-blue-400 cursor-pointer underline"
          >
            Register here
          </Link>
        </span>
      </div>
    </div>
  );
}
