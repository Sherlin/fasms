import React from "react";
var ApplicantTable = function (_a) {
    var applicants = _a.applicants, onEdit = _a.onEdit, onDelete = _a.onDelete;
    return (React.createElement("table", { border: 1, style: { width: "100%", textAlign: "left" } },
        React.createElement("thead", null,
            React.createElement("tr", null,
                React.createElement("th", null, "ID"),
                React.createElement("th", null, "Name"),
                React.createElement("th", null, "Employment Status"),
                React.createElement("th", null, "Sex"),
                React.createElement("th", null, "Date of Birth"),
                React.createElement("th", null, "Household"),
                React.createElement("th", null, "Actions"))),
        React.createElement("tbody", null, applicants.map(function (applicant) { return (React.createElement("tr", { key: applicant.id },
            React.createElement("td", null, applicant.nric),
            React.createElement("td", null, applicant.name),
            React.createElement("td", null, applicant.employment_status),
            React.createElement("td", null, applicant.sex),
            React.createElement("td", null, applicant.date_of_birth),
            React.createElement("td", null, applicant.household),
            React.createElement("td", null,
                React.createElement("button", { onClick: function () { return onEdit(applicant); } }, "Edit"),
                React.createElement("button", { onClick: function () { return onDelete(applicant.id); } }, "Delete")))); }))));
};
export default ApplicantTable;
//# sourceMappingURL=ApplicantTable.js.map