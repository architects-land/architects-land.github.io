import "./Hero.scss";

export default function Hero(props: any) {
    let clazz = "hero"
    if (props.dark) {
        clazz += " hero--dark"
    }
    if (props.min) {
        clazz += " hero--min"
    }
    return (
        <>
            <div class={clazz} style={`background-image: url("${props.image}")`}>
                <div class="hero__text__container big-text">
                    <div class="hero__text">
                        <h1>{props.title}</h1>
                        <p>
                            {props.description}
                        </p>
                    </div>
                </div>
                <div class="hero__space"></div>
            </div>
        </>
    )
}