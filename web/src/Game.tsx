import {useState} from "react";
import {Link} from "react-router-dom";
import useWebSocket from "react-use-websocket";

export const Game = () => {
    const [board, setBoard] = useState<null | Array<Array<number>>>();
    const [won, setWon] = useState<boolean>(false);
    const [lost, setLost] = useState<boolean>(false);
    const [draw, setDraw] = useState<boolean>(false);

    const { sendJsonMessage } = useWebSocket("ws://localhost:8080/game", {
        onMessage: (messageEvent) => {
            let message = JSON.parse(messageEvent.data)
            if (message.type === "board") {
                setBoard(message.data)
            }
            if (message.type === "winner") {
                switch (String.fromCharCode(message.data.symbol)) {
                    case "P":
                        setWon(true);
                        break
                    case "A":
                        setLost(true);
                }
            }
            if (message.type === "draw") {
                setDraw(true);
            }
        }
    })

    const sendAction = (row: number, column: number) => {
        sendJsonMessage({row, column});
    }

    if (won) {
        return (
            <>
                <p>Gratulation Sie haben gewonnen!</p>
                <Link to={"/"}>
                    <button>Zur Startseite</button>
                </Link>
                <Link to={"/game"} reloadDocument>
                    <button>Neue Runde</button>
                </Link>
            </>
        );
    }
    if (lost) {
        return (
            <>
                <p>Sie haben leider verloren!</p>
                <Link to={"/"}>
                    <button>Zur Startseite</button>
                </Link>
                <Link to={"/game"} reloadDocument>
                    <button>Neue Runde</button>
                </Link>
            </>
        );
    }
    if (draw) {
        return (
            <>
                <p>Unentschieden!</p>
                <Link to={"/"}>
                    <button>Zur Startseite</button>
                </Link>
                <Link to={"/game"} reloadDocument>
                    <button>Neue Runde</button>
                </Link>
            </>
        );
    }

    if (board) {
        return (
            <>
                {
                    board.map((row, rowIndex) => {
                            return (
                                <div key={`row-${rowIndex}`}>
                                    {row.map((column, columnIndex) => {
                                        let symbol;
                                        switch (String.fromCharCode(column)) {
                                            case "P":
                                                symbol = "X";
                                                break
                                            case "A":
                                                symbol = "O";
                                                break
                                            default:
                                                symbol = "-";
                                                break
                                        }
                                        return <button key={`row-${rowIndex}-column-${columnIndex}`}
                                                       onClick={() => sendAction(rowIndex, columnIndex)}>{symbol}</button>
                                    })}
                                </div>
                            );
                        }
                    )
                }
                <Link to={"/"}>
                    <button>Zur Startseite</button>
                </Link>
                <Link to={"/game"} reloadDocument>
                    <button>Neustart</button>
                </Link>
            </>
        )
    } else {
        return <p>Connecting...</p>
    }
}

export default Game
