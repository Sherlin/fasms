import React, { useState, useEffect } from "react";
var ApplicantForm = function (_a) {
    var initialData = _a.initialData, onSubmit = _a.onSubmit, onCancel = _a.onCancel;
    var _b = useState(""), name = _b[0], setName = _b[1];
    var _c = useState(""), employmentStatus = _c[0], setEmploymentStatus = _c[1];
    var _d = useState(""), employmentStatus = _d[0], setEmploymentStatus = _d[1];
    var _e = useState(""), sex = _e[0], setSex = _e[1];
    var _f = useState(""), dateOfBirth = _f[0], setDateOfBirth = _f[1];
    var _g = useState(""), household = _g[0], setHousehold = _g[1];
    useEffect(function () {
        if (initialData) {
            setNRIC(initialData.nric);
            setName(initialData.name);
            setEmploymentStatus(initialData.employment_status);
            setSex(initialData.sex);
            setDateOfBirth(initialData.date_of_birth);
            setHousehold(initialData.household);
        }
    }, [initialData]);
    var handleSubmit = function (e) {
        e.preventDefault();
        onSubmit({ nric: nric, name: name,employment_status: employmentStatus, sex: sex, date_of_birth: dateOfBirth, household: household });
    };
    return (React.createElement("form", { onSubmit: handleSubmit },
        React.createElement("div", null,
            React.createElement("label", null, "Name: "),
            React.createElement("input", { value: name, onChange: function (e) { return setName(e.target.value); }, required: true })),
        React.createElement("div", null,
            React.createElement("label", null, "Employment Status: "),
            React.createElement("input", { value: nric, onChange: function (e) { return setNRIC(e.target.value); }, required: true })),
        React.createElement("div", null,
            React.createElement("label", null, "Employment Status: "),
            React.createElement("input", { value: employmentStatus, onChange: function (e) { return setEmploymentStatus(e.target.value); }, required: true })),
        React.createElement("div", null,
            React.createElement("label", null, "Sex: "),
            React.createElement("input", { value: sex, onChange: function (e) { return setSex(e.target.value); }, required: true })),
        React.createElement("div", null,
            React.createElement("label", null, "Date of Birth: "),
            React.createElement("input", { type: "date", value: dateOfBirth, onChange: function (e) { return setDateOfBirth(e.target.value); }, required: true })),
        React.createElement("div", null,
            React.createElement("label", null, "Household: "),
            React.createElement("input", { value: household, onChange: function (e) { return setHousehold(e.target.value); }, required: true })),
        React.createElement("button", { type: "submit" }, "Submit"),
        React.createElement("button", { type: "button", onClick: onCancel }, "Cancel")));
};
export default ApplicantForm;
//# sourceMappingURL=ApplicantForm.js.map