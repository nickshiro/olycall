import { lazy } from "react";
import { Routes, Route } from "react-router";

// Pages
const AuthPage = lazy(() => import("@/pages/Auth"));
const RoomPage = lazy(() => import("@/pages/Room"));

const Main = () => {
	return (
		<Routes>
			<Route path="/auth" element={<AuthPage />} />
			<Route path="/room" element={<RoomPage />} />
		</Routes>
	);
};

export { Main };
