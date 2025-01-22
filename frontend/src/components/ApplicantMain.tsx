import React, { useEffect, useState } from "react";
import { Applicant, getApplicants, addApplicant, updateApplicant, deleteApplicant } from "../services/api";
import ApplicantTable from "./ApplicantTable";
//import ApplicantForm from "./components/ApplicantForm";
import ApplicantManager from "./ApplicantManager";
const ApplicantMain : React.FC = () => {
    const [applicants, setApplicants] = useState<Applicant[]>([]);
    const [editingApplicant, setEditingApplicant] = useState<any | null>(null);
    const [isModalOpen, setIsModalOpen] = useState(false);
    useEffect(() => {
        const fetchData = async () => {
            const data = await getApplicants();
            setApplicants(data);
        };
        fetchData();
    }, []);
    
    const handleAdd = async (applicant: Omit<Applicant, "id">) => {
        const newApplicant = await addApplicant(applicant);
        setApplicants([...applicants, newApplicant]);
        setIsModalOpen(false); // Close the modal
    };
    
    const handleEdit = async (id: string, applicant: Omit<Applicant, "id">) => {
        setEditingApplicant(applicant);
        const updatedApplicant = await updateApplicant(id, applicant);
        setApplicants(applicants.map((a) => (a.id === id ? updatedApplicant : a)));
        setEditingApplicant(null);
        setIsModalOpen(false); // Open the modal
    };
    
    const handleEditModal = (applicant: any) => {
        setEditingApplicant(applicant); // Set the selected applicant
        setIsModalOpen(true); // Open the modal
    };
    
    const closeModal = () => {
        setEditingApplicant(null); // Clear the selected applicant
        setIsModalOpen(false); // Close the modal
    };


    const openCreateApplicantModal = () => {
        setEditingApplicant(null); // Ensure no applicant is being edited
        setIsModalOpen(true); // Open modal
    };
    
    const handleDelete = async (id: string) => {
        await deleteApplicant(id);
        setApplicants(applicants.filter((a) => a.id !== id));
    };
    

    return (
        <div>
            <h3>Applicants</h3>
            <button onClick={openCreateApplicantModal}>Create Applicant</button>
            <ApplicantManager
                editingApplicant={editingApplicant}
                isModalOpen={isModalOpen}
                onCloseModal={closeModal}
                handleAdd={handleAdd}
                handleEdit={handleEdit}
   
            />
            {/*
            {editingApplicant ? (
                <ApplicantForm
                    initialData={editingApplicant}
                    onSubmit={(data) => handleEdit(editingApplicant.id, data)}
                    onCancel={() => setEditingApplicant(null)}
                />
            ) : (
                <ApplicantForm onSubmit={handleAdd} onCancel={() => {}} />
            )}
            */}
            <ApplicantTable
                applicants={applicants}
                onEdit={handleEditModal}
                onDelete={handleDelete}
            />
        </div>
    );
};

export default ApplicantMain;