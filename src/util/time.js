const calcWatchTime = (episodes, runtime) => {
    if (!episodes || !runtime) {
        return "???";
    }
    const totalSeconds = episodes * runtime;
    const days = Math.floor(totalSeconds / 60 / 60 / 24);
    const hours = Math.floor(totalSeconds / 60 / 60) % 24;
    const minutes = Math.floor(totalSeconds / 60) % 60;
    return `${days} days, ${hours} hours, ${minutes} minutes`;
};

const calRunTime = (runtime) => {
    if (!runtime) {
        return "???";
    }

    return `${Math.floor(runtime / 60)} mins`;
};

module.exports = {
    calcWatchTime,
    calRunTime,
}