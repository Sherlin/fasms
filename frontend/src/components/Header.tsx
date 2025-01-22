import React, { useState } from "react";
import ApplicantMain from "./ApplicantMain";

const Scheme = () => <div>Scheme Component</div>;
const Dependent = () => <div>Dependent Component</div>;

const Header: React.FC = () => {
    const [activeComponent, setActiveComponent] = useState<string | null>(null);

    const handleButtonClick = (component: string) => {
        setActiveComponent(component);
    };

    return (
        <div>
            <header>
                <h1> FAS Management System</h1>
                <div>
                    <button onClick={() => handleButtonClick("Applicant")}>Applicant</button>
                    <button onClick={() => handleButtonClick("Scheme")}>Scheme</button>
                    <button onClick={() => handleButtonClick("Dependent")}>Dependent</button>
                </div>
            </header>

            <div>
                {activeComponent === "Applicant" && <ApplicantMain />}
                {activeComponent === "Scheme" && <Scheme />}
                {activeComponent === "Dependent" && <Dependent />}
            </div>
        </div>
    );
};

export default Header;