import {useState} from "react";
import {Link} from "react-router-dom";

interface Statistic {
    total: number,
    win: number,
    lose: number,
    draw: number,
}

const Statistic = () => {
    const [statistic, setStatistic] = useState<null | Statistic>();
    fetch("http://localhost:8080/statistic")
        .then(res => res.json())
        .then(json => setStatistic(json))
        .catch(err => console.error(err));
    if (statistic) {
        return (
            <>
                <p>Der Computer hat bis jetzt {statistic.total} Spiel(e) gespielt.</p>
                <p>Davon hat er...</p>
                <ul>
                    <li>{statistic.win} gewonnen.</li>
                    <li>{statistic.lose} verloren.</li>
                    <li>{statistic.draw} unentschieden gespielt.</li>
                </ul>
                <Link to={"/"}>
                    <button>Zur Startseite</button>
                </Link>
            </>
        );
    } else {
        return <p>Loading...</p>
    }
}

export default Statistic;