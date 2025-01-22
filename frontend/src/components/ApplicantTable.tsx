import React from "react";
import { Applicant } from "../services/api";

interface ApplicantTableProps {
    applicants: Applicant[];
    onEdit: (applicant: any) => void;
    onDelete: (id: string) => void;
}

const ApplicantTable: React.FC<ApplicantTableProps> = ({ applicants, onEdit, onDelete }) => {
    return (
        <table border={1} style={{ width: "100%", textAlign: "left" }}>
            <thead>
                <tr>
                    <th>NRIC</th>
                    <th>Name</th>
                    <th>Employment Status</th>
                    <th>Sex</th>
                    <th>Date of Birth</th>
                    <th>Household</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {applicants.map((applicant) => (
                    <tr key={applicant.nric}>
                        <td>{applicant.nric}</td>
                        <td>{applicant.name}</td>
                        <td>{applicant.employment_status}</td>
                        <td>{applicant.sex}</td>
                        <td>{applicant.date_of_birth}</td>
                        <td>{applicant.household}</td>
                        <td>
                            <button onClick={() => onEdit(applicant)}>Edit</button>
                            <button onClick={() => onDelete(applicant.id)}>Delete</button>
                        </td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
};

export default ApplicantTable;