import { SubmitHandler } from "react-hook-form";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { FaEye } from "react-icons/fa";
import { FaEyeSlash } from "react-icons/fa";

type RegisterType = {
  name: string;
  email: string;
  username: string;
  password: string;
};

export default function RegisterForm() {
  const [hidePassword, setHidePassword] = useState(true);
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<RegisterType>();
  const onSubmit: SubmitHandler<RegisterType> = (data) => {
    console.log(data);
  };

  const passwordEye = hidePassword ? (
    <FaEye className="cursor-pointer" onClick={() => setHidePassword(false)} />
  ) : (
    <FaEyeSlash
      className="cursor-pointer"
      onClick={() => setHidePassword(true)}
    />
  );
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
      <div>
        <span className="label-text">Name</span>
        <input
          type="text"
          placeholder="Your Name"
          className="input input-bordered w-full max-w-xs"
          {...register("name", { required: true })}
        />
      </div>
      <div>
        <span className="label-text">Email</span>
        <input
          type="text"
          placeholder="Your Email"
          className="input input-bordered w-full max-w-xs"
          {...register("email", { required: true })}
        />
      </div>
      <div>
        <span className="label-text">Username</span>
        <input
          type="text"
          placeholder="Your Username"
          className="input input-bordered w-full max-w-xs"
          {...register("username", { required: true })}
        />
      </div>
      <div>
        <span className="label-text">Password</span>
        <label className="input input-bordered flex items-center gap-2">
          <input
            type={`${hidePassword ? "password" : "text"}`}
            placeholder="Your Password"
            className="grow h-full w-full"
            {...register("password", { required: true })}
          />
          {passwordEye}
        </label>
      </div>
      <button className="btn btn-primary">Register</button>
    </form>
  );
}
