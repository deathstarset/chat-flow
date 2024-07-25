import { SubmitHandler, useForm } from "react-hook-form";
import { FaEye } from "react-icons/fa";
import { FaEyeSlash } from "react-icons/fa";
import { useState } from "react";
type LoginType = {
  username: string;
  password: string;
};

export default function LoginForm() {
  const [hidePassword, setHidePassword] = useState(true);
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<LoginType>();
  const onSubmit: SubmitHandler<LoginType> = (data) => {
    console.log(data);
  };
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
      <div>
        <span className="label-text">Username</span>
        <input
          type="text"
          placeholder="Your Username"
          className="input input-bordered w-full max-w-xs"
          {...(register("username"), { required: true })}
        />
      </div>
      <div>
        <span className="label-text">Password</span>
        <label className="input input-bordered flex items-center gap-2">
          <input
            type={`${hidePassword ? "password" : "text"}`}
            placeholder="Your Password"
            className="grow h-full w-full"
            {...(register("password"), { required: true })}
          />
          {hidePassword ? (
            <FaEye
              className="cursor-pointer"
              onClick={() => setHidePassword(false)}
            />
          ) : (
            <FaEyeSlash
              className="cursor-pointer"
              onClick={() => setHidePassword(true)}
            />
          )}
        </label>
      </div>
      <button className="btn btn-primary">Log In</button>
    </form>
  );
}
