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
import {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {AiOutlineHome} from "react-icons/all";
import {Link} from "react-router-dom";

interface Statistic {
    total: number,
    win: number,
    lose: number,
    draw: number,
}

const Statistic = () => {
    const {t} = useTranslation();
    const [statistic, setStatistic] = useState<null | Statistic>(null);
    useEffect(() => {
        fetch(`${import.meta.env.VITE_BACKEND}/statistic`)
            .then(res => res.json())
            .then(json => setStatistic(json))
            .catch(err => console.error(err));
    }, []);
    if (statistic) {
        return (
            <>
                <Card className={"card"}>
                    <CardHeader>
                        <Heading>{t("title.statistic")}</Heading>
                    </CardHeader>
                    <CardBody>
                        <Text pb={2}>{t("statistic.introduction", {
                            rounds: statistic.total,
                            roundWord: statistic.total === 1 ? t("statistic.round") : t("statistic.rounds")
                        })}</Text>
                        <Stat>
                            <StatLabel>{t("statistic.roundsWon")}</StatLabel>
                            <StatNumber>{statistic.win}</StatNumber>
                        </Stat>
                        <Stat>
                            <StatLabel>{t("statistic.roundsLost")}</StatLabel>
                            <StatNumber>{statistic.lose}</StatNumber>
                        </Stat>
                        <Stat>
                            <StatLabel>{t("statistic.roundsDraw")}</StatLabel>
                            <StatNumber>{statistic.draw}</StatNumber>
                        </Stat>
                    </CardBody>
                    <CardFooter>
                        <Link to={"/"}>
                            <Button leftIcon={<AiOutlineHome/>}>{t("button.toHome")}</Button>
                        </Link>
                    </CardFooter>
                </Card>
            </>
        );
    } else {
        return (
            <Center width={"100vw"} height={"100svh"}>
                <Spinner size={"xl"}/>
            </Center>
        );
    }
}

export default Statistic;