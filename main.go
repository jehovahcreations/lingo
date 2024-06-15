package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
type Aplica struct{
	SUCCESS	bool `json:"success"`
		STATUSCODE	int `json:"statusCode"`
		STATUS	string `json:"status"`
		MESSAGE	string `json:"message"`
		Data	struct {
			
				
				APPLICATIONID string `json:"applicationId"`
				
		} `json:"data"`
}
	type Response struct {
		SUCCESS    bool    `json:"success"`
		STATUSCODE int `json:"statusCode"`
		STATUS string `json:"status"`
		MESSAGE string `json:"message"`
		Data    struct {
			
				USERID   string `json:"userId"`
				TOKEN string `json:"token"`
			
		} `json:"data"`
	}
	type Login struct {
		SUCCESS    bool    `json:"success"`
		STATUSCODE int `json:"statusCode"`
		STATUS string `json:"status"`
		MESSAGE string `json:"message"`
		Data    struct {
			
				USERID   string `json:"userId"`
			
		} `json:"data"`
	}
	type Otp struct{
		SUCCESS	bool `json:"success"`
		STATUSCODE	int `json:"statusCode"`
		STATUS	string `json:"status"`
		MESSAGE	string `json:"message"`
		Data	struct {
			
				USERID   string `json:"userId"`
				TOKEN string `json:"token"`
				APPLICATIONID string `json:"applicationId"`
				CONSUMERID string `json:"consumerId"`
				ISAPPLICATIONREFERENCE bool `json:"isApplicationReference"`
				ISAPPLICATIONLOCK bool `json:"isApplicationLock"`
				REFERENCEID int `json:"referenceId"` 
				LEADID string `json:"leadID"`
		} `json:"data"`
	}
	type FinalStep struct{
		SUCCESS	bool `json:"success"`
		STATUSCODE	int `json:"statusCode"`
		STATUS	string `json:"status"`
		MESSAGE	string `json:"message"`
		Data	struct {
			
				USERID   string `json:"userId"`
				TOKEN string `json:"token"`
				APPLICATIONID string `json:"applicationId"`
				CONSUMERID string `json:"consumerId"`
				ISAPPLICATIONREFERENCE bool `json:"isApplicationReference"`
				ISAPPLICATIONLOCK bool `json:"isApplicationLock"`
				REFERENCEID int `json:"referenceId"` 
				LEADID string `json:"leadID"`
		} `json:"data"`
	}
	type Finalupdate struct{
		SUCCESS bool `json:"success"`
		STATUSCODE int `json:"statusCode"`
		STATUS string `json:"status"`
		MESSAGE string `json:"message"`
		Data    struct {
			USERID   string `json:"userId"`
			CONSUMERID string `json:"consumerId"`
			ISAPPLICATIONREFERENCE bool `json:"isApplicationReference"`
			ISAPPLICATIONLOCK bool `json:"isApplicationLock"`
			APPLICATIONID string `json:"applicationId"`
			REFERENCEID int `json:"referenceId"` 
			USERTYPE int `json:"userType"`
		}`json:"data"`}

	type Profile struct{
		SUCCESS bool `json:"success"`
		STATUSCODE int `json:"statusCode"`
		STATUS string `json:"status"`
		MESSAGE string `json:"message"`
		Data    struct {
			USERID   string `json:"userId"`
			CONSUMERID string `json:"consumerId"`
			ISAPPLICATIONREFERENCE bool `json:"isApplicationReference"`
			ISAPPLICATIONLOCK bool `json:"isApplicationLock"`
			APPLICATIONID string `json:"applicationId"`
			REFERENCEID int `json:"referenceId"` 
			USERTYPE int `json:"userType"`
		}`json:"data"`}

		type AppDetail struct{
			
    DATA struct{
        APPLICATION struct {
            ADDRESS1 string `json:"address1"`
            ADDRESS2 string `json:"address2"`
            APPLICATIONID string `json:"applicationId"`
            APPLICATIONNAME string `json:"applicationName"`
            APPLICATIONNUMBER string `json:"applicationNumber"`
            APPLICATIONSTATUSID int `json:"applicationStatusID"`
            APPLICATIONSUBMITTEDDATE string `json:"applicationSubmittedDate"`
            CATEGORYID int `json:"categoryID"`
            CATEGORYNAME string `json:"categoryName"`
            CIRCLEID int `json:"circleID"`
            CIRCLENAME string `json:"circleName"`
            CITYNAME string `json:"cityName"`
            CONSUMER_VENDOR_ID string `json:"consumer_vendor_id"`
            CONSUMERCONNECTIONCATEGORYID int `json:"consumerConnectionCategoryID"`
            CONSUMERCONNECTIONCATEGORYNAME string `json:"consumerConnectionCategoryName"`
            CONSUMERID string `json:"consumerID"`
            CONSUMERNUMBER string `json:"consumerNumber"`
            CSCVLENUMBER *int `json:"cscVLENumber"`
            CURRENT_STEP_ID int `json:"statusCode"`
            CURRENTSTEPID int `json:"current_step_id"`
            DISCOMID int `json:"discomID"`
            DISCOMNAME string `json:"discomName"`
            DISTRICTID int `json:"districtID"`
            DISTRICTNAME string `json:"districtName"`
            DIVISIONID int `json:"divisionID"`
            DIVISIONNAME string `json:"divisionName"`
            EMAIL string `json:"email"`
            EXISTINGCAPACITY string `json:"existingCapacity"`
            FEASIBILITYREGISTRATIONURL string `json:"feasibilityRegistrationUrl"`
			HASCSCVLE int `json:"hasCSCVLE"`
            INSTALLATIONFORMURL string `json:"installationFormUrl"`
            ISAPPLICATIONLOCK int `json:"isApplicationLock"`
			ISDELETED int `json:"isDeleted"`
            LATITUDE *string `json:"latitude"`
            LEADID int `json:"leadId"`
            LONGITUDE *string `json:"longitude"`
            MOBILENUMBER string `json:"mobileNumber"`
            NAME string `json:"name"`
            NETMETERFEESURL string `json:"netMeterFeesUrl"`
            PINCODE int `json:"pinCode"`
            PVCAPACITY string `json:"pvCapacity"`
            REGISTRATIONNUMBER string `json:"registrationNumber"`
            RESUBMITTEDDATE string `json:"reSubmittedDate"`
            SANCTIONLOAD string `json:"sanctionLoad"`
            SCHEMEID int `json:"schemeID"`
            SCHEMENAME string `json:"schemeName"`
            STATEID int `json:"stateID"`
            STATENAME string `json:"stateName"`
            STATUSNAME string `json:"statusName"`
            SUBDIVISIONID int `json:"subDivisionID"`
            SUBDIVISIONNAME string `json:"subDivisionName"`
            SUBMITEDDATE string `json:"submitedDate"`
            SUBSIDYCOST string `json:"subsidyCost"`
            VENDOREMAIL string `json:"vendorEmail"`
            VENDORMOBILE string `json:"vendorMobile"`
            VENDORNAME string `json:"vendorName"`
            VILLAGE string `json:"village"`
        }`json:"application"`
        DOCUMENTS struct {}
    }`json:"data"`
    MESSAGE string `json:"message"`
    STATUS string `json:"status"`
    STATUSCODE int `json:"statusCode"`
    SUCCESS string `json:"success"`
		}

		type UpdateD struct{
	       Address1 string `json:"address1"`
           Address2 *string `json:"address2"`
           ApplicationId string `json:"applicationId"`
           ApplicationNumber string `json:"applicationNumber"`
           ApplicationStatusID int `json:"applicationStatusID"`
           CategoryID int `json:"categoryID"`
           CircleID int `json:"circleID"`
           CityName string `json:"cityName"`
           ConsumerConnectionCategoryID int `json:"consumerConnectionCategoryID"`
           ConsumerNumber string `json:"consumerNumber"`
           CscVLENumber *int `json:"cscVLENumber"`
           CurrentStepID int `json:"currentStepID"`
           DiscomID int `json:"discomID"`
           DiscomName string `json:"discomName"`
           DistrictID int `json:"districtID"`
           DistrictName string `json:"districtName"`
           DivisionID int `json:"divisionID"`
           Email string `json:"email"`
           ExistingCapacity string `json:"existingCapacity"`
           ExistingSolarCapacity string `json:"existingSolarCapacity"`
           HasCSCVLE int `json:"hasCSCVLE"`
           Latitude *string `json:"latitude"`
           Longitude *string `json:"longitude"`
           MobileNumber string `json:"mobileNumber"`
           Name string `json:"name"`
           PinCode int `json:"pinCode"`
           PvCapacity string `json:"pvCapacity"`
           SanctionLoad string `json:"sanctionLoad"`
           StateID int `json:"stateID"`
           StateName string `json:"stateName"`
           SubDivisionID int `json:"subDivisionID"`
       	   Village string `json:"village"`
		}

		type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserName     string             `bson:"userName"`
	Password     string             `bson:"password"`
	Token      int                `bson:"token"`
	CscAuth     string             `bson:"cscAuth"`
	StateId     string             `bson:"stateId"`
	District      string                `bson:"district"`
	Discom     string             `bson:"discom"`
	Consumer     string             `bson:"consumer"`
	Ack      string                `bson:"ack"`
}

type Auto struct{
Id        primitive.ObjectID `bson:"_id"`
Name string `bson:"name"`
Consumer string `bson:"consumer"`
Mobile string `bson:"mobile"`
Status int    `bson:"status"`
City string `bson:"city"`
Village string `bson:"village"`
Pin int    `bson:"pin"`
Add string `bson:"add"`
Circle int    `bson:"circle"`
Discom int    `bson:"discom"`
District int    `bson:"district"`
Division int    `bson:"division"`
State int    `bson:"state"`
Sub int    `bson:"sub"`
}
		
	var cooki string	
	var isDone bool
    var autos []Auto
	// var circleID int = 535
	// var Address1 string ="Neeral"
	// var city string ="Raigad"
	// var village string ="Neeral"
	 var cscId int = 523360500014
	// var districtID int = 950
	// var division int = 1913
	// var name string ="vibhute"
	// var email string ="bnet@ghgd.gdf"
	// var pinCode int = 421503
	// var stateId int =27
	// var SubDiv int =5272
	

func close(client *mongo.Client, ctx context.Context,cancel context.CancelFunc){
    defer cancel()
    defer func(){
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
	}
 
func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error) {
	credential := options.Credential{
    Username: "myMongoAdmin",
    Password: "Jesmysav@123",
}
    ctx, cancel := context.WithTimeout(context.Background(), 
                                       30 * time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
    return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error{
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}
func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
    collection := client.Database(dataBase).Collection(col)
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}
func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{})(result *mongo.Cursor, err error) {
    collection := client.Database(dataBase).Collection(col)
    result, err = collection.Find(ctx, query, 
                                  options.Find().SetProjection(field))
    return
}
func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{})(result *mongo.UpdateResult, err error) {
    collection := client.Database(dataBase).Collection(col)
    result, err = collection.UpdateOne(ctx, filter, update)
    return
}

var myId = "665b23e83aac3b5d1a8f67f3"

var MongoDb = "mongodb://172.214.106.123:27017/?authSource=admin"

func main(){
   GetHomePage()
	}

func GetHomePage(){

	client1, ctx, cancel, err := connect(MongoDb)
    if err != nil {
        panic(err)
    }
    defer close(client1, ctx, cancel)
    ping(client1, ctx)
	var filter, option interface{}

   filter = bson.M{"status":0}
//   filter = bson.D{
//     {"status", 0},
//     {"$or", bson.A{
//         bson.D{{"s", 30}},
//         bson.D{{"a", 10}},
//     }},
// }
	
    cursor, err :=query(client1, ctx, "surya", "autos", filter, option)
	 if err != nil {
        panic(err)
    }
	if err := cursor.All(ctx, &autos); err != nil {
        panic(err)
    }
	fmt.Println(autos)
	
	for i := 0; i < len(autos); i++{ 
       sendOtp(autos[i].Mobile,i)
    } 
		fmt.Scanln()
	
	}
func sendOtp(mobile string,i int){
			 client, ctx, cancel, err := connect(MongoDb)
    			if err != nil {
    			    panic(err)
    			}
    			defer close(client, ctx, cancel)
			ping(client,ctx)
	if len(autos)>0{
		isDone=false
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		clientX := &http.Client{Transport: tr}
		var data = strings.NewReader(`{"mobileNumber":`+mobile+`,"email":""}`)
		fmt.Println("Processing ...............\n\n")
		req, err := http.NewRequest("POST", "https://www.pmsuryaghar.gov.in/api/auth/consumerLogin", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Host", "www.pmsuryaghar.gov.in")
		req.Header.Set("Content-Length", "40")
		req.Header.Set("Sec-Ch-Ua", `"Not-A.Brand";v="99", "Chromium";v="124"`)
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.60 Safari/537.36")
		req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
		req.Header.Set("Origin", "https://www.pmsuryaghar.gov.in")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Referer", "https://www.pmsuryaghar.gov.in/consumerLogin")
		req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Set("Priority", "u=1, i")
		resp, err := clientX.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		// bodyText, err := io.ReadAll(resp.Body)
		// fmt.Printf("%s\n", bodyText)
		login := &Login{}
		json.NewDecoder(resp.Body).Decode(login)
		fmt.Println(login)
		fmt.Println(login.MESSAGE + "\n\n")
		fmt.Println("Please wait .......................\n\n")
		if err != nil {
			log.Fatal(err)
		}
		spl1 := strings.Split(resp.Header["Set-Cookie"][0], "AWSALB=")
		spl2 := strings.Split(spl1[1], ";")
		out, err := json.Marshal(login)
    if err != nil {
        panic (err)
    }

    fmt.Println(string(out))
	if login.SUCCESS{
		otpVerify(mobile,i,spl2[0],login.Data.USERID)
	
			}else{
				fmt.Println("OTP Request Failed")
			}
 		}else{
			fmt.Print("No data Available")
		}
}			
func otpVerify(mobile string,i int,spl2 string,userid string){
	time.Sleep(3*time.Second)
	 ctx1 := context.Background()
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		clientX := &http.Client{Transport: tr}
		var otp ="497816"
		var data1 = strings.NewReader(`{"mobileNumber":`+mobile+`,"otp":"`+otp+`","userId":"`+userid+`"}`)
		req1, err := http.NewRequest("POST", "https://www.pmsuryaghar.gov.in/api/auth/verifyConsumerMobileOtp", data1)
		if err != nil {
			log.Fatal(err)
		}
		req1.Header.Set("Host", "www.pmsuryaghar.gov.in")
		req1.Header.Set("Content-Length", "89")
		req1.Header.Set("Sec-Ch-Ua", `"Not-A.Brand";v="99", "Chromium";v="124"`)
		req1.Header.Set("Accept", "application/json, text/plain, */*")
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req1.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.60 Safari/537.36")
		req1.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
		req1.Header.Set("Origin", "https://www.pmsuryaghar.gov.in")
		req1.Header.Set("Sec-Fetch-Site", "same-origin")
		req1.Header.Set("Sec-Fetch-Mode", "cors")
		req1.Header.Set("Sec-Fetch-Dest", "empty")
		req1.Header.Set("Referer", "https://www.pmsuryaghar.gov.in/consumerLogin")
		// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req1.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
		req1.Header.Set("Priority", "u=1, i")
		req1.Header.Set("Cookie", "AWSALB="+spl2+"; AWSALBCORS="+spl2)
		ctx3, cancel := context.WithDeadline(ctx1, time.Now().Add(3 * time.Second))
       defer cancel()

        req1.WithContext(ctx3)
		resp1, err := clientX.Do(req1)
		if err != nil {
			log.Fatal(err)
		}
		defer resp1.Body.Close()
		// bodyText2, err := io.ReadAll(resp1.Body)
		// fmt.Printf("%s\n",bodyText2)
		otpResponse := &Otp{}
		json.NewDecoder(resp1.Body).Decode(otpResponse)
		fmt.Println(otpResponse)
		
		fmt.Println(otpResponse.MESSAGE+"\n")
		
		fmt.Println("Please wait .......................\n\n")
		fmt.Println("generating token .......................\n\n" +otpResponse.Data.APPLICATIONID)
	out, err := json.Marshal(otpResponse)
    if err != nil {
        panic (err)
    }

    fmt.Println(string(out))

		if otpResponse.SUCCESS {
			if len(otpResponse.Data.APPLICATIONID)>0{
           		fmt.Println("\n token :"+otpResponse.Data.TOKEN+"\n")
				fmt.Println("\n otp app :"+otpResponse.Data.APPLICATIONID+"\n")
				 client, ctx, cancel, err := connect(MongoDb)
    			if err != nil {
    			    panic(err)
    			}
    			defer close(client, ctx, cancel)
				ping(client,ctx)
                filter := bson.M{"_id":autos[i].Id}
                update := bson.D{{Key: "$set", Value: bson.M{"token": otpResponse.Data.TOKEN,"status":1}, }}
                result, err := UpdateOne(client, ctx, "surya", "autos", filter, update)
				
                if err != nil {
                   panic(err)
                }
                fmt.Println(result.ModifiedCount)
				appDat(otpResponse.Data.TOKEN,otpResponse.Data.APPLICATIONID,i)

			}else{
				fmt.Println("Lead id"+otpResponse.Data.LEADID)
				var data2 = strings.NewReader(`{"leadID":"`+otpResponse.Data.LEADID+`"}`)
	req2, err := http.NewRequest("POST", "https://www.pmsuryaghar.gov.in/api/application/processApplicationDetails", data2)
	if err != nil {
		log.Fatal(err)
	}
	req2.Header.Set("Host", "www.pmsuryaghar.gov.in")
	req2.Header.Set("Content-Length", "49")
	req2.Header.Set("Sec-Ch-Ua", `"Chromium";v="125", "Not.A/Brand";v="24"`)
	req2.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req2.Header.Set("Authorization", "Bearer "+otpResponse.Data.TOKEN)
	req2.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.112 Safari/537.36")
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Accept", "application/json, text/plain, */*")
	req2.Header.Set("X-Access-Token", "Bearer "+otpResponse.Data.TOKEN)
	req2.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req2.Header.Set("Origin", "https://www.pmsuryaghar.gov.in")
	req2.Header.Set("Sec-Fetch-Site", "same-origin")
	req2.Header.Set("Sec-Fetch-Mode", "cors")
	req2.Header.Set("Sec-Fetch-Dest", "empty")
	req2.Header.Set("Referer", "https://www.pmsuryaghar.gov.in/consumerLogin")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req2.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req2.Header.Set("Priority", "u=1, i")
	req2.Header.Set("Cookie", "AWSALB=6eaINljWXGW5D7jgZEIcp35+FpC6ZIr8I9LQhA93UScGjN9jsETrldf0jGubItlvSeSkReG8I+1kw31qfkKq8gaAf3sOPgjSbgAVaKEZOjdvYymZMIDVNI/KvtGcq1oPVT2bb9RaPY26um05df8tkbXODlXD03D8RLEaiZfzAl6qp5qjnTc7jGNCD3w4mg==; AWSALBCORS=6eaINljWXGW5D7jgZEIcp35+FpC6ZIr8I9LQhA93UScGjN9jsETrldf0jGubItlvSeSkReG8I+1kw31qfkKq8gaAf3sOPgjSbgAVaKEZOjdvYymZMIDVNI/KvtGcq1oPVT2bb9RaPY26um05df8tkbXODlXD03D8RLEaiZfzAl6qp5qjnTc7jGNCD3w4mg==")
	ctx2, cancel := context.WithDeadline(ctx1, time.Now().Add(3 * time.Second))
       defer cancel()

        req2.WithContext(ctx2)
	resp2, err := clientX.Do(req2)
		if err != nil {
			log.Fatal(err)
		}
		defer resp2.Body.Close()
    // bodyText, err := io.ReadAll(resp2.Body)
	// fmt.Println(string(bodyText))
		apppli := &Aplica{}
		json.NewDecoder(resp2.Body).Decode(apppli)
		// if err != nil {
		// log.Fatal(err)
		// }
		out, err := json.Marshal(apppli)
    if err != nil {
        panic (err)
    }

    fmt.Println(string(out))
		fmt.Println(apppli)
	if apppli.SUCCESS {
		otpResponse.Data.APPLICATIONID = apppli.Data.APPLICATIONID
		fmt.Println("\n token :"+otpResponse.Data.TOKEN+"\n")
		fmt.Println("applicationId :"+apppli.Data.APPLICATIONID+"\n")
		fmt.Print("*********Results******* \n\n")
		fmt.Print("Insert Id : ")
		fmt.Print(autos[i].Id)
		 client, ctx, cancel, err := connect(MongoDb)
    			if err != nil {
    			    panic(err)
    			}
    			defer close(client, ctx, cancel)
			ping(client,ctx)
		filter1 := bson.M{"_id":autos[i].Id}
    	update1 := bson.D{{Key: "$set", Value: bson.M{"token": otpResponse.Data.TOKEN,"status":1}, }}
    	result1, err := UpdateOne(client, ctx, "surya", "autos", filter1, update1)
		appDat(otpResponse.Data.TOKEN,apppli.Data.APPLICATIONID,i)
    if err != nil {
        panic(err)
    }
		fmt.Println(result1.ModifiedCount)

	
	}else{
fmt.Print("Application process Failed")
	}
}
	
}	
	}

func appDat(token string,apppId string,i int){
	
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	client1 := &http.Client{Transport: tr}
	var data = strings.NewReader(`{"applicationId":"`+apppId+`"}`)
	//fmt.Println(data)
	req, err := http.NewRequest("POST", "https://www.pmsuryaghar.gov.in/api/application/getApplicationByIdDetail", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.pmsuryaghar.gov.in")
	req.Header.Set("Content-Length", "50")
	req.Header.Set("Sec-Ch-Ua", `"Chromium";v="125", "Not.A/Brand";v="24"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.60 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("X-Access-Token", "Bearer "+token)
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Origin", "https://www.pmsuryaghar.gov.in")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.pmsuryaghar.gov.in/consumer/applicationForm")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Cookie", "AWSALB=id+wEyzspIuLTSWz6Le2Weg61/FDF9Nq0zRAbIrDSDDSCdeA/t1JjstmArFRulGa4WgdgBuJxQwdoKfUq0WS/PDIBiH6fUd1OE30mJiCbzUz/7M0IVSHkc+BLUwb3gLQm6sEh/Iqjfir9nl5szYmvRfOmKEqH8x3mLhf9yKsKO1EhBlRYOGKnBCSua/USA==; AWSALBCORS=id+wEyzspIuLTSWz6Le2Weg61/FDF9Nq0zRAbIrDSDDSCdeA/t1JjstmArFRulGa4WgdgBuJxQwdoKfUq0WS/PDIBiH6fUd1OE30mJiCbzUz/7M0IVSHkc+BLUwb3gLQm6sEh/Iqjfir9nl5szYmvRfOmKEqH8x3mLhf9yKsKO1EhBlRYOGKnBCSua/USA==")
	resp, err := client1.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	appData := &AppDetail{}
	defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(appData)
		out, err := json.Marshal(appData)
    if err != nil {
        panic (err)
    }

    fmt.Println(string(out))
		fmt.Println("appdata :")
		fmt.Println(appData)
		if appData.DATA.APPLICATION.STATUSNAME == "Feasibility Approved with Applied Capacity"{
client, ctx, cancel, err := connect(MongoDb)
    			if err != nil {
    			    panic(err)
    			}
    			defer close(client, ctx, cancel)
			ping(client,ctx)
	filter22 := bson.M{"_id":autos[i].Id}
	 update22 := bson.D{{Key: "$set", Value: bson.M{"status":4} }}
    result22, err := UpdateOne(client, ctx, "surya", "autos", filter22, update22)
	if err != nil {
        panic(err)
    }
    // fmt.Println("update single document")
    fmt.Println(result22.ModifiedCount)
	fmt.Print("Completed Sucessfully")
		}else{
	if err != nil {
		log.Fatal(err)
	}
	updateData := &UpdateD{}
	updateData.ApplicationId = appData.DATA.APPLICATION.APPLICATIONID
	updateData.ApplicationNumber = appData.DATA.APPLICATION.APPLICATIONNUMBER
	updateData.ApplicationStatusID =1
	updateData.CategoryID=1
	updateData.ConsumerConnectionCategoryID=0
	updateData.Address2=nil
	updateData.Address1 = autos[i].Add
	updateData.CircleID = autos[i].Circle
	updateData.CityName = autos[i].City
	updateData.ConsumerNumber = appData.DATA.APPLICATION.CONSUMERNUMBER
	updateData.CscVLENumber = &cscId
	updateData.HasCSCVLE = 1
	updateData.CurrentStepID=1
	updateData.DiscomID=autos[i].Discom
	updateData.DiscomName =appData.DATA.APPLICATION.DISCOMNAME
	updateData.DistrictID = autos[i].District
	updateData.DistrictName = appData.DATA.APPLICATION.DISTRICTNAME
	updateData.DivisionID = autos[i].Division
	updateData.ExistingCapacity ="0.00"
	updateData.MobileNumber = appData.DATA.APPLICATION.MOBILENUMBER
	updateData.Name = autos[i].Name
	updateData.Email = ""
	updateData.PinCode = autos[i].Pin
	updateData.Longitude = nil
	updateData.Latitude = nil
	updateData.PvCapacity = "1.00"
	updateData.SanctionLoad ="1.00"
	updateData.StateID = autos[i].State
	updateData.StateName = appData.DATA.APPLICATION.STATENAME
	updateData.SubDivisionID = autos[i].Sub
	updateData.Village = autos[i].Village
	println(updateData)
	b, err := json.Marshal(updateData)
	var data1 = strings.NewReader(string(b))
	req1, err := http.NewRequest("POST", "https://www.pmsuryaghar.gov.in/api/application/updateApplication", data1)
	if err != nil {
		log.Fatal(err)
	}
	req1.Header.Set("Host", "www.pmsuryaghar.gov.in")
	req1.Header.Set("Content-Length", "741")
	req1.Header.Set("Sec-Ch-Ua", `"Chromium";v="125", "Not.A/Brand";v="24"`)
	req1.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req1.Header.Set("Authorization", "Bearer "+token)
	req1.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.60 Safari/537.36")
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("Accept", "application/json, text/plain, */*")
	req1.Header.Set("X-Access-Token", "Bearer "+token)
	req1.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req1.Header.Set("Origin", "https://www.pmsuryaghar.gov.in")
	req1.Header.Set("Sec-Fetch-Site", "same-origin")
	req1.Header.Set("Sec-Fetch-Mode", "cors")
	req1.Header.Set("Sec-Fetch-Dest", "empty")
	req1.Header.Set("Referer", "https://www.pmsuryaghar.gov.in/consumer/applicationForm")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req1.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req1.Header.Set("Priority", "u=1, i")
	req1.Header.Set("Cookie", "AWSALB=ZbXg91KP1LCegRRbu4YCilfKNumkvx0404QYMVBo7hdDq081sZz6s6qDjlzdpy/IAjLXuDnw/+gfcbpynwck5oGkNvhBuJcYBFXScrpugPmK0tdlF4QZqZfrEnjHqMZfGvZEi9nnaU09erctUVM6wrNqILqLIhwnxImbD4lgqUlHYwx0/Dxw7MZ7rLp7Ug==; AWSALBCORS=ZbXg91KP1LCegRRbu4YCilfKNumkvx0404QYMVBo7hdDq081sZz6s6qDjlzdpy/IAjLXuDnw/+gfcbpynwck5oGkNvhBuJcYBFXScrpugPmK0tdlF4QZqZfrEnjHqMZfGvZEi9nnaU09erctUVM6wrNqILqLIhwnxImbD4lgqUlHYwx0/Dxw7MZ7rLp7Ug==")
	resp1, err := client1.Do(req1)
	if err != nil {
		log.Fatal(err)
	}
	defer resp1.Body.Close()
	//bodyText, err := io.ReadAll(resp1.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	finalstep := FinalStep{}
	if finalstep.SUCCESS{
	 client, ctx, cancel, err := connect(MongoDb)
    			if err != nil {
    			    panic(err)
    			}
    			defer close(client, ctx, cancel)
			ping(client,ctx)
	filter2 := bson.M{"_id":autos[i].Id}
    update2 := bson.D{{Key: "$set", Value: bson.M{"status":2}, }}
    result2, err := UpdateOne(client, ctx, "surya", "autos", filter2, update2)
	fmt.Println(result2.ModifiedCount)
    if err != nil {
        panic(err)
    }
}
    
	fmt.Print("Please wait for the Final updation.....")
	var data2 = strings.NewReader(`{"applicationId":"`+apppId+`"}`)
	req2, err := http.NewRequest("POST", "https://www.pmsuryaghar.gov.in/api/application/updateApplicationFinalStep", data2)
	if err != nil {
		log.Fatal(err)
	}
	req2.Header.Set("Host", "www.pmsuryaghar.gov.in")
	req2.Header.Set("Content-Length", "50")
	req2.Header.Set("Sec-Ch-Ua", `"Chromium";v="125", "Not.A/Brand";v="24"`)
	req2.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req2.Header.Set("Authorization", "Bearer "+token)
	req2.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.112 Safari/537.36")
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Accept", "application/json, text/plain, */*")
	req2.Header.Set("X-Access-Token", "Bearer "+token)
	req2.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req2.Header.Set("Origin", "https://www.pmsuryaghar.gov.in")
	req2.Header.Set("Sec-Fetch-Site", "same-origin")
	req2.Header.Set("Sec-Fetch-Mode", "cors")
	req2.Header.Set("Sec-Fetch-Dest", "empty")
	req2.Header.Set("Referer", "https://www.pmsuryaghar.gov.in/consumer/applicationForm")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req2.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req2.Header.Set("Priority", "u=1, i")
	req2.Header.Set("Cookie", "AWSALB=KBcYetYuy3CoTlPEwAbeeYGXjy1wRb1FqQkEwv2GhkeaobHFvW7bB5UJBQS1HfQCwkufd9tL2HoIfrCDRFmayMGFxLQSxGAK3AvdfUr83WgT9OgK2gBXmR66/6hKJaLXX9HOiDQjh5sBqCY7hVP52dkACBS7+58mllbU0Ti+EM55VMCJkRPvXe1tVRv9rA==; AWSALBCORS=KBcYetYuy3CoTlPEwAbeeYGXjy1wRb1FqQkEwv2GhkeaobHFvW7bB5UJBQS1HfQCwkufd9tL2HoIfrCDRFmayMGFxLQSxGAK3AvdfUr83WgT9OgK2gBXmR66/6hKJaLXX9HOiDQjh5sBqCY7hVP52dkACBS7+58mllbU0Ti+EM55VMCJkRPvXe1tVRv9rA==")
	resp2, err := client1.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()
	// bodyText2, err := io.ReadAll(resp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(bodyText2))
	updatefinal := Finalupdate{}
	if updatefinal.SUCCESS{
	 client, ctx, cancel, err := connect(MongoDb)
    			if err != nil {
    			    panic(err)
    			}
    			defer close(client, ctx, cancel)
			ping(client,ctx)
    filter13 := bson.M{"_id":autos[i].Id}
    update13 := bson.D{{Key: "$set", Value: bson.M{"status":3} }}
    result13, err := UpdateOne(client, ctx, "surya", "autos", filter13, update13)
    if err != nil {
        panic(err)
    }
    fmt.Println(result13.ModifiedCount)
	}
///final form

	var userr = &User{}
    	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	formField, err := writer.CreateFormField("ci_csrf_token")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte(""))

	formField, err = writer.CreateFormField("state")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte(userr.StateId))

	formField, err = writer.CreateFormField("district")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte(userr.District))

	formField, err = writer.CreateFormField("distribution_company")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte(userr.Discom))

	formField, err = writer.CreateFormField("constumer_number")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte(updateData.ConsumerNumber))

	formField, err = writer.CreateFormField("ack")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte(updateData.ApplicationNumber))

	formField, err = writer.CreateFormField("submit")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte("1"))

	writer.Close()

	tr1 := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client7 := &http.Client{Transport: tr1}
	req7, err := http.NewRequest("POST", "https://api.pmsuryabijliyojna.in/acknowledge", form)
	if err != nil {
		log.Fatal(err)
	}
	req7.Header.Set("Host", "api.pmsuryabijliyojna.in")
	req7.Header.Set("Content-Length", "765")
	req7.Header.Set("Cache-Control", "max-age=0")
	req7.Header.Set("Sec-Ch-Ua", `"Chromium";v="125", "Not.A/Brand";v="24"`)
	req7.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req7.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req7.Header.Set("Upgrade-Insecure-Requests", "1")
	req7.Header.Set("Origin", "https://api.pmsuryabijliyojna.in")
	req7.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.112 Safari/537.36")
	req7.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req7.Header.Set("Sec-Fetch-Site", "same-origin")
	req7.Header.Set("Sec-Fetch-Mode", "navigate")
	req7.Header.Set("Sec-Fetch-User", "?1")
	req7.Header.Set("Sec-Fetch-Dest", "document")
	req7.Header.Set("Referer", "https://api.pmsuryabijliyojna.in/")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req7.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req7.Header.Set("Priority", "u=0, i")
	req7.Header.Set("Cookie", "ci_sessions="+userr.CscAuth)
	req7.Header.Set("Content-Type", writer.FormDataContentType())
	resp7, err := client7.Do(req7)
	if err != nil {
		log.Fatal(err)
	}
	defer resp7.Body.Close()
	
	client, ctx, cancel, err := connect(MongoDb)
    			if err != nil {
    			    panic(err)
    			}
    			defer close(client, ctx, cancel)
			ping(client,ctx)
	filter22 := bson.M{"_id":autos[i].Id}
	 update22 := bson.D{{Key: "$set", Value: bson.M{"status":4} }}
    result22, err := UpdateOne(client, ctx, "surya", "autos", filter22, update22)
    if err != nil {
        panic(err)
    }
    // fmt.Println("update single document")
    fmt.Println(result22.ModifiedCount)
	fmt.Print("Completed Sucessfully")
}	
	
}


		