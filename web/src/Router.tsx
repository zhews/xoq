import {createBrowserRouter} from "react-router-dom";
import Home from "./Home";
import Game from "./Game";
import NotFound from "./NotFound";
import Statistic from "./Statistic";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Home />,
        errorElement: <NotFound/>,
    },
    {
        path: "/game",
        element: <Game />,
    },
    {
        path: "/statistic",
        element: <Statistic />,
    }
]);

export default router;
