import {useEffect, useState} from "react";
import {Link} from "react-router-dom";
import {
    Button,
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    Center,
    Heading,
    Spinner,
    Stat,
    StatLabel,
    StatNumber,
    Text
} from "@chakra-ui/react";
import { AiOutlineHome } from "react-icons/all";

interface Statistic {
    total: number,
    win: number,
    lose: number,
    draw: number,
}

const Statistic = () => {
    const [statistic, setStatistic] = useState<null | Statistic>(null);
    useEffect(() => {
        fetch("http://localhost:8080/statistic")
            .then(res => res.json())
            .then(json => setStatistic(json))
            .catch(err => console.error(err));
    });
    if (statistic) {
        return (
            <>
                <Card className={"card"}>
                    <CardHeader>
                        <Heading>Statistiken</Heading>
                    </CardHeader>
                    <CardBody>
                        <Text pb={2}>Der Computer hat bis jetzt {statistic.total} {statistic.total === 1 ? "Runde" : "Runden"} gespielt.</Text>
                        <Stat>
                            <StatLabel>Gewonnene Runden</StatLabel>
                            <StatNumber>{statistic.win}</StatNumber>
                        </Stat>
                        <Stat>
                            <StatLabel>Verlorene Runden</StatLabel>
                            <StatNumber>{statistic.lose}</StatNumber>
                        </Stat>
                        <Stat>
                            <StatLabel>Unentschiedene Runden</StatLabel>
                            <StatNumber>{statistic.draw}</StatNumber>
                        </Stat>
                    </CardBody>
                    <CardFooter>
                        <Link to={"/"}>
                            <Button leftIcon={<AiOutlineHome/>}>Zur Startseite</Button>
                        </Link>
                    </CardFooter>
                </Card>
            </>
        );
    } else {
        return (
            <Center width={"100vw"} height={"100vh"}>
                <Spinner size={"xl"}/>
            </Center>
        );
    }
}

export default Statistic;