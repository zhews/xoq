import {Link} from "react-router-dom";

export const Home = () => {
    return (
        <>
            <h1>xoq</h1>
            <p>Für unsere Interdisziplinäre Projektarbeit mit dem Überthema "Künstliche Intelligenz (KI)" haben wir uns folgende Frage gestellt.</p>
            <blockquote>Welche Macht kann eine KI durch verabeitete Daten aus dem Internet erlangen?</blockquote>
            <p>Zum einschätzen dieser Aussage wurde dieses Tic-Tac-Toe Spiel entwickelt.</p>
            <p>Es lernt durch die Spielzüge aller Spieler und verbessert sich fortlaufend.</p>
            <p>Im Hintergrund verwendet es ein Q-Lernen Algorithmus.</p>
            <Link to={"/game"}>
                <button>Runde Starten</button>
            </Link>
            <Link to={"/statistic"}>
                <button>Statistiken</button>
            </Link>
        </>
    );
}

export default Home
