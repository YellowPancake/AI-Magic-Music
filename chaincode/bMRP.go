package main

import (
	"encoding/json"
   	 "fmt"
	//"strconv"
  //  "strings"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)
type bMRPChainCode struct{	
}

type Patient struct{
    PatientName string `json:PatientName` //患者姓名
    PatientGender int `json:PatientGender` //患者性别
    PatientAge string `json:PatientAge` //患者年龄
    PatientNationality string `json:PatientNationality` //国籍
    PatientIDType string `json:PatientIDType` //证件类型
    PatientIDNumber string `json:PatientIDNumber` //证件编号
    PatientTelephone string `json:PatientTelephone` //电话号码
	PatientAddress string `json:PatientAddress` //住址
}

type Doctor struct{
    DoctorName string `json:DoctorName` //医生姓名
    DoctorID string `json:DoctorID` //医生编号
    DoctorHospitalName string `json:DoctorHospitalName` //医院名称
	DoctorHosptialID string `json:DoctorHosptialID` //医院编号
}

type MedicalRecord struct{
	MRID string `json:MRID`//病历编号
	MRAdmissionDate string `json:MRAdmissionDate` //就诊日期、时间
    MRDischargeDate string  `json:MRDischargeDate`//出院日期、时间
    MRPaymentType string  `json:MRPaymentType`//付款方式：自费、医保
    MRPatientID string `json:MRPatientID`
    MRDoctors string `json:MRDoctors`
    MRDiagnosis string `json:MRDiagnosis` //诊断内容（包括影像资料）
}


func(a *bMRPChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}


func(a *bMRPChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn,args := stub.GetFunctionAndParameters()
	if fn == "AddNewMR"{
		return a.AddNewMR(stub,args)
	}else if fn=="GetMRByID"{
        return a.GetMRByID(stub,args)
    }

    return shim.Error("Recevied unkown function invocation : "+fn)
}

func(a *bMRPChainCode) AddNewMR(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var err error
	var newRecord MedicalRecord
	//检查参数个数是否正确
	if len(args)!=7{
		return shim.Error("Incorrect number of arguments.")
	}
	newRecord.MRID=args[0]
	newRecord.MRAdmissionDate=args[1]
    newRecord.MRDischargeDate=args[2]
    newRecord.MRPaymentType=args[3]
    newRecord.MRPatientID=args[4]
    newRecord.MRDoctors =args[5]
	newRecord.MRDiagnosis =args[6]
	
	ProInfosJSONasBytes,err := json.Marshal(newRecord)
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(newRecord.MRID,ProInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}

func(a *bMRPChainCode) GetMRByID(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    MRID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(MRID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    //var foodProInfo ProInfo
    var medicalRecord MedicalRecord

    for resultsIterator.HasNext(){
        var _medicalRecord MedicalRecord
        //var FoodInfos FoodInfo
        response,err :=resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&_medicalRecord)
        if _medicalRecord.MRID != ""{
            medicalRecord = _medicalRecord
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(medicalRecord)   
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func main(){
     err := shim.Start(new(bMRPChainCode))
     if err != nil {
         fmt.Printf("Error starting bMRP chaincode: %s ",err)
     }
}

