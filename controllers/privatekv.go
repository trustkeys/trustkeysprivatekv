package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/trustkeys/trustkeysprivatekv/common/util"
	"github.com/trustkeys/trustkeysprivatekv/models"
)

var (
	tkPrivateModel models.TrustkeysPrivatekvModelIf
	enable_getsig bool
)

// Operations about Public Key-value store
type PrivateKVController struct {
	beego.Controller
}

// Set model for controller
func SetPrivateModel(aModel models.TrustkeysPrivatekvModelIf, en_getsig bool) {
	enable_getsig = en_getsig
	tkPrivateModel = aModel
}

var GetMessage = util.GetMessage

// @Title PutItem
// @Description Put key-value to cloud
// @Param	pubKey		path 	string	true		"Public key in hex"
// @Param	appID		path 	string	true		"App ID"
// @Param	key		query  	string	true		"The Key"
// @Param	val		query  	string	true		"The Value"
// @Param 	sig		query 	string 	true 	"signature of a message = "TrustKeys:" + pubKey + appID + Key + Value "
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /putitem/:appID/:pubKey [post]
func (o *PrivateKVController) PutItem(pubKey, appID string) {

	// pubKey := o.GetString("pubKey")
	sig := o.GetString("sig")

	fmt.Println("PutItem ", o)

	//Todo: check sessionID

	appID = o.GetString(":appID")
	key := o.GetString("key")
	val := o.GetString("val")
	fmt.Printf("pubKey: %s, appID: %s, object: %s : %s, sig: %s \n", pubKey, appID, key, val, sig)

	if util.CheckSignature(pubKey, GetMessage(pubKey, appID, key, val), sig) {
		var aMap = map[string]string{"Key": key, "errCode": "0", "desc": "success"}

		if tkPrivateModel != nil {
			ok, oldVal, transID := tkPrivateModel.Put(appID, pubKey, key, val)
			if ok {
				aMap["oldValue"] = oldVal
				aMap["transactionID"] = transID
			} else {
				aMap["errCode"] = "-2" // not ok
				aMap["desc"] = "backend error"
			}
		}
		o.Data["json"] = aMap

	} else {
		o.Data["json"] = map[string]string{"Key": key, "errCode": "-1", "desc": "Invalid signature!"}
	}

	o.ServeJSON()
}

// @Title PutSafeItem
// @Description Put key-value to cloud extra timestamp
// @Param	pubKey		path 	string	true		"Public key in hex"
// @Param	appID		path 	string	true		"App ID"
// @Param	key		query  	string	true		"The Key"
// @Param	val		query  	string	true		"The Value"
// @Param	timeStamp		query  	string	true		"Timestamp"
// @Param 	sig		query 	string 	true 	"signature of a message = "TrustKeys:" + pubKey + appID + Key + Value + timestamp"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /putsafeitem/:appID/:pubKey [post]
func (o *PrivateKVController) PutSafeItem(pubKey, appID string) {

	// pubKey := o.GetString("pubKey")
	sig := o.GetString("sig")

	fmt.Println("PutItem ", o)

	//Todo: check sessionID

	appID = o.GetString(":appID")
	key := o.GetString("key")
	val := o.GetString("val")
	timeStamp := o.GetString("timeStamp")
	fmt.Printf("pubKey: %s, appID: %s, object: %s : %s, timestamp : %s , sig: %s \n", pubKey, appID, key, val, sig)

	if util.CheckSignature(pubKey, GetMessage(pubKey, appID, key, val+timeStamp), sig) {
		var aMap = map[string]string{"Key": key, "errCode": "0", "desc": "success"}

		if tkPrivateModel != nil {
			ok, oldVal, transID := tkPrivateModel.Put(appID, pubKey, key, val)
			if ok {
				aMap["oldValue"] = oldVal
				aMap["transactionID"] = transID
			} else {
				aMap["errCode"] = "-2" // not ok
				aMap["desc"] = "backend error"
			}
		}
		o.Data["json"] = aMap

	} else {
		o.Data["json"] = map[string]string{"Key": key, "errCode": "-1", "desc": "Invalid signature!"}
	}

	o.ServeJSON()
}


// @Title GetItem
// @Description find key-value by key with check sig
// @Param	pubKey		query 	string	true		"Public Key of a user"
// @Param	appID		query 	string	true		"appID"
// @Param	key		query 	string	true		"the key of kv you want to get"
// @Param	sig		query 	string	true		"signature of a message = "TrustKeys:" + pubKey + appID + Key"
// @Success 200 {object} models.KVObject
// @Failure 403 : empty object
// @router /get [get]
func (o *PrivateKVController) GetItem() {
	pubKey := o.GetString("pubKey")
	appID := o.GetString("appID")
	key := o.GetString("key") //o.Ctx.Input.Param(":key")
	sig := o.GetString("sig")
	fmt.Printf("pubKey: %s, appID: %s, key: %s", pubKey, appID, key)

	if util.CheckSignature(pubKey, GetMessage(pubKey, appID, key, ""), sig) && enable_getsig {
		var aMap = map[string]string{"Key": pubKey}
		if tkPrivateModel != nil {
			ok, value, lastestTrans := tkPrivateModel.Get(appID, pubKey, key)
			if ok {
				// aMap["errCode"] = "0";
				// aMap["value"] = value;
				// aMap["lastestTransactionID"] = lastestTrans
				o.Data["json"] = &models.KVObject{
					Key:           key,
					Value:         value,
					TransactionID: lastestTrans,
				}
				o.ServeJSON()
				return
			} else {
				aMap["errCode"] = "-1"
				aMap["desc"] = "Model error or empty data"
			}

		}
		o.Data["json"] = aMap
	} else {
		o.Data["json"] = map[string]string{"Key": key, "errCode": "-1", "desc": "Invalid signature!"}
	}

	o.ServeJSON()
}


// @Title GetSliceFrom
// @Description find key-value by key with check sig
// @Param	pubKey		query 	string	true		"Public Key of a user"
// @Param	appID		query 	string	true		"appID"
// @Param	fromKey		query 	string	true		"the key of kv you want to get"
// @Param	maxNum		query 	int	true		"Maximum number of items to get"
// @Param	sig		query 	string	true		"signature of a message = "TrustKeys:" + pubKey + appID + fromKey + maxNum"
// @Success 200 {array} []models.KVObject
// @Failure 403 : empty object
// @router /GetSliceFrom/:appID/:pubKey [get]
func (o *PrivateKVController) GetSliceFrom(pubKey, appID string) {
	fromKey := o.GetString("fromKey") //o.Ctx.Input.Param(":key")
	maxNum, _ := o.GetInt32("maxNum")
	sig := o.GetString("sig")
	fmt.Println("appID: ", appID, " pubKey: ", pubKey, " fromKey: ", fromKey)

	if util.CheckSignature(pubKey, GetMessage(pubKey, appID, fromKey, strconv.Itoa(int(maxNum))), sig) && enable_getsig {
		if tkPrivateModel != nil {
			kvs, err := tkPrivateModel.GetSlice(appID, pubKey, fromKey, maxNum)
			if err == nil {
				o.Data["json"] = kvs
			} else {
				o.Data["json"] = map[string]string{"errCode": "-1", "desc": "error get slice"}
			}

		}
		// o.Data["json"] = aMap;

	} else {
		o.Data["json"] = map[string]string{"Key": pubKey, "errCode": "-1", "desc": "Invalid signature!"}
	}

	o.ServeJSON()
}
