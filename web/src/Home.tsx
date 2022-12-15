import {Link} from "react-router-dom";
import {Button, Card, CardBody, CardFooter, CardHeader, Heading, Stack, Text} from "@chakra-ui/react";
import {BiStats, BsPlay} from "react-icons/all";

export const Home = () => {
    return (
        <>
            <Card className={"card"}>
                <CardHeader>
                    <Heading>xoq</Heading>
                </CardHeader>
                <CardBody>
                    <Stack>
                        <Text>Für unsere Interdisziplinäre Projektarbeit mit dem Überthema "Künstliche Intelligenz (KI)"
                            haben
                            wir uns folgende Frage gestellt.</Text>
                        <Text as={"cite"}>Zu welchem Zweck werden Daten im Internet gesammelt?</Text>
                        <Text>Zum einschätzen dieser Aussage wurde dieses Tic-Tac-Toe Spiel entwickelt.</Text>
                        <Text>Es lernt durch die Spielzüge aller Spieler und verbessert sich fortlaufend.</Text>
                        <Text>Im Hintergrund verwendet es ein Q-Lernen Algorithmus.</Text>

                    </Stack>
                </CardBody>
                <CardFooter justify={"space-between"}>
                    <Link to={"/game"}>
                        <Button leftIcon={<BsPlay/>}>Runde Starten</Button>
                    </Link>
                    <Link to={"/statistic"}>
                        <Button leftIcon={<BiStats/>}>Statistiken</Button>
                    </Link>
                </CardFooter>
            </Card>
        </>
    );
}

export default Home
