import RegisterForm from "./RegisterForm";
import { Link } from "react-router-dom";

export default function Register() {
  return (
    <div className="card bg-base-100 w-96 shadow-xl">
      <div className="card-body flex flex-col gap-4">
        <h2 className="card-title">Register</h2>
        <RegisterForm />
        <span className="text-center text-sm">
          Already have an account ?{" "}
          <Link
            to="/auth/login"
            className="text-blue-400 cursor-pointer underline"
          >
            Login here
          </Link>
        </span>
      </div>
    </div>
  );
}
