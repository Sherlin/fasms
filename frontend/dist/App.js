var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g = Object.create((typeof Iterator === "function" ? Iterator : Object).prototype);
    return g.next = verb(0), g["throw"] = verb(1), g["return"] = verb(2), typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
import React, { useEffect, useState } from "react";
import { getApplicants, addApplicant, updateApplicant, deleteApplicant } from "./services/api";
import ApplicantTable from "./components/ApplicantTable";
import ApplicantForm from "./components/ApplicantForm";
var App = function () {
    var _a = useState([]), applicants = _a[0], setApplicants = _a[1];
    var _b = useState(null), editingApplicant = _b[0], setEditingApplicant = _b[1];
    useEffect(function () {
        var fetchData = function () { return __awaiter(void 0, void 0, void 0, function () {
            var data;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        console.log("fetching data");
                        return [4, getApplicants()];
                    case 1:
                        data = _a.sent();
                        console.log("app");
                        console.log(data);
                        setApplicants(data);
                        return [2];
                }
            });
        }); };
        fetchData();
    }, []);
    var handleAdd = function (applicant) { return __awaiter(void 0, void 0, void 0, function () {
        var newApplicant;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4, addApplicant(applicant)];
                case 1:
                    newApplicant = _a.sent();
                    setApplicants(__spreadArray(__spreadArray([], applicants, true), [newApplicant], false));
                    return [2];
            }
        });
    }); };
    var handleEdit = function (id, applicant) { return __awaiter(void 0, void 0, void 0, function () {
        var updatedApplicant;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4, updateApplicant(id, applicant)];
                case 1:
                    updatedApplicant = _a.sent();
                    setApplicants(applicants.map(function (a) { return (a.id === id ? updatedApplicant : a); }));
                    setEditingApplicant(null);
                    return [2];
            }
        });
    }); };
    var handleDelete = function (id) { return __awaiter(void 0, void 0, void 0, function () {
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4, deleteApplicant(id)];
                case 1:
                    _a.sent();
                    setApplicants(applicants.filter(function (a) { return a.id !== id; }));
                    return [2];
            }
        });
    }); };
    return (React.createElement("div", null,
        React.createElement("h1", null, "Applicant Management"),
        editingApplicant ? (React.createElement(ApplicantForm, { initialData: editingApplicant, onSubmit: function (data) { return handleEdit(editingApplicant.id, data); }, onCancel: function () { return setEditingApplicant(null); } })) : (React.createElement(ApplicantForm, { onSubmit: handleAdd, onCancel: function () { } })),
        React.createElement(ApplicantTable, { applicants: applicants, onEdit: setEditingApplicant, onDelete: handleDelete })));
};
export default App;
//# sourceMappingURL=App.js.map