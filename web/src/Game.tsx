import {
    Alert,
    AlertIcon,
    Button,
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    Center,
    Heading,
    SimpleGrid,
    Spinner
} from "@chakra-ui/react";
import {useState} from "react";
import {useTranslation} from "react-i18next";
import {AiOutlineHome, VscDebugRestart} from "react-icons/all";
import {Link} from "react-router-dom";
import useWebSocket from "react-use-websocket";

export const Game = () => {
    const {t} = useTranslation();
    const [board, setBoard] = useState<null | Array<Array<number>>>();
    const [done, setDone] = useState<boolean>(false);
    const [won, setWon] = useState<boolean>(false);
    const [lost, setLost] = useState<boolean>(false);
    const [draw, setDraw] = useState<boolean>(false);

    const {sendJsonMessage} = useWebSocket(`${import.meta.env.VITE_BACKEND_WEBSOCKET}/game`, {
        onMessage: (messageEvent) => {
            let message = JSON.parse(messageEvent.data)
            if (message.type === "board") {
                setBoard(message.data)
            }
            if (message.type === "winner") {
                switch (String.fromCharCode(message.data.symbol)) {
                    case "P":
                        setWon(true);
                        setDone(true)
                        break
                    case "A":
                        setLost(true);
                        setDone(true)
                        break
                }
            }
            if (message.type === "draw") {
                setDraw(true);
                setDone(true);
            }
        }
    })


    const sendAction = (row: number, column: number) => {
        sendJsonMessage({row, column});
    }

    if (board) {
        return (
            <>
                <Card className={"card"}>
                    <CardHeader>
                        <Heading>{t("title.game")}</Heading>
                    </CardHeader>
                    <CardBody>
                        <Center mb={5}>
                            <SimpleGrid columns={3} spacing={5}>
                                {
                                    board.map((row, rowIndex) => {
                                            return (
                                                <>
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
                                                                symbol = "";
                                                                break
                                                        }
                                                        return (
                                                            <Button height="100px" width={"100px"}
                                                                    key={`row-${rowIndex}-column-${columnIndex}`}
                                                                    disabled={done}
                                                                    onClick={() => sendAction(rowIndex, columnIndex)}>{symbol}</Button>
                                                        );
                                                    })}
                                                </>
                                            );
                                        }
                                    )
                                }
                            </SimpleGrid>
                        </Center>
                        {done && won && <Alert status={"success"}><AlertIcon/>{t("game.won")}</Alert>}
                        {done && lost && <Alert status={"error"}><AlertIcon/>{t("game.lost")}</Alert>}
                        {done && draw && <Alert status={"info"}><AlertIcon/>{t("game.draw")}</Alert>}
                    </CardBody>
                    <CardFooter justify={"space-between"}>
                        <Link to={"/"}>
                            <Button leftIcon={<AiOutlineHome/>}>{t("button.toHome")}</Button>
                        </Link>
                        <Link to={"/game"} reloadDocument>
                            <Button leftIcon={<VscDebugRestart/>}>{t("button.restart")}</Button>
                        </Link>
                    </CardFooter>
                </Card>
            </>
        )
    } else {
        return (
            <Center width={"100vw"} height={"100vh"}>
                <Spinner size={"xl"}/>
            </Center>
        )
    }
}

export default Game
