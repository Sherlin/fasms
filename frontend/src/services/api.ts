import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:8080/api",
});

export interface Applicant {
    id: string;
    nric: string;
    name: string;
    employment_status: string;
    sex: string;
    date_of_birth: string;
    household: string;
}

// Fetch all applicants
export const getApplicants = async (): Promise<Applicant[]> => {
    const response = await api.get("/applicants/min");
    console.log(response.data);
    return response.data;
};

// Add a new applicant
export const addApplicant = async (applicant: Omit<Applicant, "id">): Promise<Applicant> => {
    const response = await api.post("/applicants", applicant);
    return response.data;
};

// Update an existing applicant
export const updateApplicant = async (id: string, applicant: Omit<Applicant, "id">): Promise<Applicant> => {
    const response = await api.put(`/applicants/${id}`, applicant);
    return response.data;
};

// Delete an applicant
export const deleteApplicant = async (id: string): Promise<void> => {
    await api.delete(`/applicants/${id}`);
};