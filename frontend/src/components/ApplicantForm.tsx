import React, { useState, useEffect } from "react";
import { Applicant } from "../services/api";

interface ApplicantFormProps {
    initialData?: Applicant | null;
    onSubmit: (applicant: Omit<Applicant, "id">) => void;
    onCancel: () => void;
}

const ApplicantForm: React.FC<ApplicantFormProps> = ({ initialData, onSubmit, onCancel }) => {
    const [nric, setNRIC] = useState("");
    const [name, setName] = useState("");
    const [employmentStatus, setEmploymentStatus] = useState("");
    const [sex, setSex] = useState("");
    const [dateOfBirth, setDateOfBirth] = useState("");
    const [household, setHousehold] = useState("");

    useEffect(() => {
        if (initialData) {
            setNRIC(initialData.nric);
            setName(initialData.name);
            setEmploymentStatus(initialData.employment_status);
            setSex(initialData.sex);
            setDateOfBirth(initialData.date_of_birth);
            setHousehold(initialData.household);
        }
    }, [initialData]);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        onSubmit({ nric: nric ,name :name ,employment_status: employmentStatus, sex, date_of_birth: dateOfBirth, household });
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>NRIC: </label>
                <input value={nric} onChange={(e) => setNRIC(e.target.value)} required />
            </div>
            <div>
                <label>Name: </label>
                <input value={name} onChange={(e) => setName(e.target.value)} required />
            </div>

            <div>
                <label>Employment Status: </label>
                <input value={employmentStatus} onChange={(e) => setEmploymentStatus(e.target.value)} required />
            </div>
            <div>
                <label>Sex: </label>
                <input value={sex} onChange={(e) => setSex(e.target.value)} required />
            </div>
            <div>
                <label>Date of Birth: </label>
                <input type="date" value={dateOfBirth} onChange={(e) => setDateOfBirth(e.target.value)} required />
            </div>
            <div>
                <label>Household: </label>
                <input value={household} onChange={(e) => setHousehold(e.target.value)} required />
            </div>
            <button type="submit">Submit</button>
            <button type="button" onClick={onCancel}>Cancel</button>
        </form>
    );
};

export default ApplicantForm;