import "./Hero.scss";

export default function Hero(props: any) {
    return (
        <>
            <div class="hero" style={`background: url("/${props.image}") fixed center no-repeat`}>
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