import { BrowserRouter, Routes, Route } from "react-router-dom";
import Auth from "./routes/auth/Auth";
import Login from "./routes/auth/login/Login";
import Register from "./routes/auth/register/Register";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/auth" element={<Auth />}>
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
