import React, { useEffect, useState } from "react";
import { Applicant, getApplicants, addApplicant, updateApplicant, deleteApplicant } from "./services/api";
import ApplicantTable from "./components/ApplicantTable";
import ApplicantForm from "./components/ApplicantForm";

const App: React.FC = () => {
    const [applicants, setApplicants] = useState<Applicant[]>([]);
    const [editingApplicant, setEditingApplicant] = useState<Applicant | null>(null);

    useEffect(() => {
        const fetchData = async () => {
           console.log("fetching data");
            const data = await getApplicants();
            console.log("app");
            console.log(data);
            setApplicants(data);
        };
        fetchData();
    }, []);

    const handleAdd = async (applicant: Omit<Applicant, "id">) => {
        const newApplicant = await addApplicant(applicant);
        setApplicants([...applicants, newApplicant]);
    };

    const handleEdit = async (id: string, applicant: Omit<Applicant, "id">) => {
        const updatedApplicant = await updateApplicant(id, applicant);
        setApplicants(applicants.map((a) => (a.id === id ? updatedApplicant : a)));
        setEditingApplicant(null);
    };

    const handleDelete = async (id: string) => {
        await deleteApplicant(id);
        setApplicants(applicants.filter((a) => a.id !== id));
    };

    return (
        <div>
            <h1>Applicant Management</h1>
            {editingApplicant ? (
                <ApplicantForm
                    initialData={editingApplicant}
                    onSubmit={(data) => handleEdit(editingApplicant.id, data)}
                    onCancel={() => setEditingApplicant(null)}
                />
            ) : (
                <ApplicantForm onSubmit={handleAdd} onCancel={() => {}} />
            )}
            <ApplicantTable
                applicants={applicants}
                onEdit={setEditingApplicant}
                onDelete={handleDelete}
            />
        </div>
    );
};

export default App;