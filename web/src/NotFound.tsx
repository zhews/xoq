import {Link} from "react-router-dom";

const NotFound = () => {
    return (
        <>
            <h1>Diese Seite konnte nicht gefunden werden!</h1>
            <Link to={"/"}>
                <button>Zur Startseite</button>
            </Link>
        </>
    )
}

export default NotFound;