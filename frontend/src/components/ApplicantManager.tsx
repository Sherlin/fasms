import React, { useState } from "react";
import ApplicantForm from "./ApplicantForm"; // Assuming ApplicantForm is a separate component
import { Applicant } from "../services/api";
function ApplicantManager({
    editingApplicant,
    isModalOpen,
    onCloseModal,
    handleAdd,
    handleEdit
}: {
    editingApplicant: any | null;
    isModalOpen: boolean;
    onCloseModal: () => void;
    handleAdd: (applicant: Omit<Applicant, "id">) => void;
    handleEdit: (id: string, applicant: Omit<Applicant, "id">) => void;
}) {
    //const [isModalOpen, setIsModalOpen] = useState(false);
    //const [editingApplicant, setEditingApplicant] = useState<any | null>(null);
    /*
    const handleAdd = (data: any) => {
        console.log("Adding Applicant:", data);
        onCloseModal(); // Close modal after adding
    };
    

    const handleEdit = (id: string, data: any) => {
        console.log("Editing Applicant:", id, data);
        onCloseModal(); // Close modal after editing
    };
    */

    return (
        <>
            {isModalOpen && (
                <div className="modal">
                    <div className="modal-content">
                        <button className="close-button" onClick={onCloseModal}>
                            &times;
                        </button>
                        {editingApplicant ? (
                            <ApplicantForm
                                initialData={editingApplicant}
                                onSubmit={(data) => handleEdit(editingApplicant.id, data)}
                                onCancel={onCloseModal}
                            />
                        ) : (
                            <ApplicantForm onSubmit={handleAdd} onCancel={onCloseModal} />
                        )}
                    </div>
                </div>
            )}
        </>
    );
}

export default ApplicantManager;