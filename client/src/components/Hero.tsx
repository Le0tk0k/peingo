import React, { FC } from "react";
import icon from '../icon.png';

const Hero: FC = () => (
    <div>
        <img alt="icon" src={icon} />
        <a href="https://github.com/Le0tk0k" rel="noopener noreferrer" target="_blank">GitHub</a>
        <a href="https://github.com/Le0tk0k" rel="noopener noreferrer" target="_blank">Twitter</a>
    </div>
);

export default Hero;