package database

import (
	"context"
	"database/sql"
	"exporter/internal/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
	CreateUser(user *models.User) error
	FindUser(user *models.User, email string)
	FindFile(file *models.File)
	ModifyUserPass(user *models.User, passwordHash string)
	SaveAcct(acct *models.Acct) error
	DeleteAcct(acct *models.Acct) error
	SaveAcctBank(acctBank *models.AcctBank) error
	DeleteAcctBank(acctBank *models.AcctBank) error
	FindAcct(accts *[]models.Acct)
	FindAcctBank(acctBanks *[]models.AcctBank)
	FindAcctBankById(acctBanks *[]models.AcctBank, acctId uint)
	FindAcctBankByAccName(acctBanks *[]models.AcctBank, accName string)
	FindAcctBankByAccNum(acctBanks *[]models.AcctBank, accNum string)
	FindAcctBankByCurrency(acctBanks *[]models.AcctBank, currency string)
	FindAcctBankByBankName(acctBanks *[]models.AcctBank, bankName string)
	FzzFindAcctBankByBankName(acctBanks *[]models.AcctBank, bankName string)
	FindAcctBankBySwiftCode(acctBanks *[]models.AcctBank, swiftCode string)
	FindAcctBankByBankAddr(acctBanks *[]models.AcctBank, bankAddr string)
	FzzFindAcctBankByBankAddr(acctBanks *[]models.AcctBank, bankAddr string)
	FzzFindAcctBankByNotes(acctBanks *[]models.AcctBank, notes string)
	FindAcctBankByNotes(acctBanks *[]models.AcctBank, notes string)
	FindAcctByAcctCode(accts *[]models.Acct, acctCode string)
	FindAcctByAcctAbbr(accts *[]models.Acct, acctAbbr string)
	FindAcctByEtyAbbr(accts *[]models.Acct, EtyAbbr string)
	FindAcctByAcctName(accts *[]models.Acct, acctName string)
	FindAcctByAcctAddr(accts *[]models.Acct, acctAddr string)
	FzzFindAcctByAcctAddr(accts *[]models.Acct, acctAddr string)
	FindAcctByNation(accts *[]models.Acct, nation string)
	FindAcctByTaxType(accts *[]models.Acct, taxType string)
	FindAcctByTaxCode(accts *[]models.Acct, taxCode string)
	FindAcctByPhoneNum(accts *[]models.Acct, phoneNum string)
	FzzFindAcctByPhoneNum(accts *[]models.Acct, phoneNum string)
	FindAcctByEmail(accts *[]models.Acct, email string)
	FzzFindAcctByEmail(accts *[]models.Acct, email string)
	FindAcctByWebsite(accts *[]models.Acct, website string)
	FzzFindAcctByWebsite(accts *[]models.Acct, website string)
	FindAcctByRegDate(accts *[]models.Acct, regDate string)
	FzzFindAcctByRegDate(accts *[]models.Acct, regDate string)
	FindAcctByNotes(accts *[]models.Acct, notes string)
	FzzFindAcctByNotes(accts *[]models.Acct, notes string)

	CreateFile(file *models.File) error
	// FindFile(file *models.File, MD5 string)

	// Save functions
	SaveMercType(mercType *models.MercType) error
	SaveSuprType(suprType *models.SuprType) error
	SavePrdtType(prdtType *models.PrdtType) error
	SaveFoodAddType(foodAddType *models.FoodAddType) error
	SaveFeedAddType(feedAddType *models.FeedAddType) error
	SaveUnitMeas(unitMeas *models.UnitMeas) error
	SavePackType(packType *models.PackType) error
	SaveConType(conType *models.ConType) error
	SaveCurrency(currency *models.Currency) error
	SaveTradeTerm(tradeTerm *models.TradeTerm) error
	SaveNation(nation *models.Nation) error
	SavePort(port *models.Port) error
	SaveTaxType(taxType *models.TaxType) error
	SaveBrandType(brandType *models.BrandType) error
	SaveDegree(degree *models.Degree) error
	SaveDept(dept *models.Dept) error
	SavePost(post *models.Post) error
	SaveQualStd(qualStd *models.QualStd) error
	SaveInvLoc(invLoc *models.InvLoc) error
	SaveDocReq(docReq *models.DocReq) error
	SavePayMth(payMth *models.PayMth) error
	SavePayLimit(payLimit *models.PayLimit) error
	SaveFinaDocStatus(finaDocStatus *models.FinaDocStatus) error
	SaveFinaDocType(finaDocType *models.FinaDocType) error
	SaveExpType(expType *models.ExpType) error
	SaveRates(rates *models.Rates) error
	SaveBussOrderSta(bussOrderSta *models.BussOrderSta) error

	// Find functions
	FindMercType(mercTypes *[]models.MercType)
	FindSuprType(suprTypes *[]models.SuprType)
	FindPrdtType(prdtTypes *[]models.PrdtType)
	FindFoodAddType(foodAddTypes *[]models.FoodAddType)
	FindFeedAddType(feedAddTypes *[]models.FeedAddType)
	FindUnitMeas(unitMeas *[]models.UnitMeas)
	FindPackType(packTypes *[]models.PackType)
	FindConType(conTypes *[]models.ConType)
	FindCurrency(currencies *[]models.Currency)
	FindTradeTerm(tradeTerms *[]models.TradeTerm)
	FindNation(nations *[]models.Nation)
	FindPort(ports *[]models.Port)
	FindTaxType(taxTypes *[]models.TaxType)
	FindBrandType(brandTypes *[]models.BrandType)
	FindDegree(degrees *[]models.Degree)
	FindDept(depts *[]models.Dept)
	FindPost(posts *[]models.Post)
	FindQualStd(qualStds *[]models.QualStd)
	FindInvLoc(invLocs *[]models.InvLoc)
	FindDocReq(docReqs *[]models.DocReq)
	FindPayMth(payMths *[]models.PayMth)
	FindPayLimit(payLimits *[]models.PayLimit)
	FindFinaDocStatus(finaDocStatuses *[]models.FinaDocStatus)
	FindFinaDocType(finaDocTypes *[]models.FinaDocType)
	FindExpType(expTypes *[]models.ExpType)
	FindRates(rates *[]models.Rates)
	FindBussOrderSta(bussOrderStas *[]models.BussOrderSta)

	// Delete functions
	DeleteMercType(mercType *models.MercType) error
	DeleteSuprType(suprType *models.SuprType) error
	DeletePrdtType(prdtType *models.PrdtType) error
	DeleteFoodAddType(foodAddType *models.FoodAddType) error
	DeleteFeedAddType(feedAddType *models.FeedAddType) error
	DeleteUnitMeas(unitMeas *models.UnitMeas) error
	DeletePackType(packType *models.PackType) error
	DeleteConType(conType *models.ConType) error
	DeleteCurrency(currency *models.Currency) error
	DeleteTradeTerm(tradeTerm *models.TradeTerm) error
	DeleteNation(nation *models.Nation) error
	DeletePort(port *models.Port) error
	DeleteTaxType(taxType *models.TaxType) error
	DeleteBrandType(brandType *models.BrandType) error
	DeleteDegree(degree *models.Degree) error
	DeleteDept(dept *models.Dept) error
	DeletePost(post *models.Post) error
	DeleteQualStd(qualStd *models.QualStd) error
	DeleteInvLoc(invLoc *models.InvLoc) error
	DeleteDocReq(docReq *models.DocReq) error
	DeletePayMth(payMth *models.PayMth) error
	DeletePayLimit(payLimit *models.PayLimit) error
	DeleteFinaDocStatus(finaDocStatus *models.FinaDocStatus) error
	DeleteFinaDocType(finaDocType *models.FinaDocType) error
	DeleteExpType(expType *models.ExpType) error
	DeleteRates(rates *models.Rates) error
	DeleteBussOrderSta(bussOrderSta *models.BussOrderSta) error
	// Merchant 相关接口
	FindMerchantByMercCode(merchants *[]models.Merchant, mercCode string)
	FindMerchantByMercAbbr(merchants *[]models.Merchant, mercAbbr string)
	FindMerchantByShortMerc(merchants *[]models.Merchant, shortMerc string)
	FindMerchantByMerc(merchants *[]models.Merchant, merc string)
	FindMerchantByEngName(merchants *[]models.Merchant, engName string)
	FindMerchantByAddress(merchants *[]models.Merchant, address string)
	FindMerchantByNation(merchants *[]models.Merchant, nation string)
	FindMerchantByPhoneNum(merchants *[]models.Merchant, phoneNum string)
	FindMerchantByEmail(merchants *[]models.Merchant, email string)
	FindMerchantByFax(merchants *[]models.Merchant, fax string)
	FindMerchantByWebsite(merchants *[]models.Merchant, website string)
	FindMerchantByTaxType(merchants *[]models.Merchant, taxType string)
	FindMerchantByTaxCode(merchants *[]models.Merchant, taxCode string)
	FindMerchantByMercType(merchants *[]models.Merchant, mercType string)
	FindMerchantByRegCap(merchants *[]models.Merchant, regCap string)
	FindMerchantByNotes(merchants *[]models.Merchant, notes string)
	DeleteMerchant(merchant *models.Merchant) error
	SaveMerchant(merchant *models.Merchant) error
	FindMerchant(merchant *[]models.Merchant)
	FzzFindMerchantByEngName(merchant *[]models.Merchant, EngName string)
	FindMerchantByMercId(merchant *[]models.Merchant, MercId uint)
	FzzFindMerchantByMercAbbr(merchant *[]models.Merchant, MercAbbr string)
	FzzFindMerchantByShortMerc(merchant *[]models.Merchant, ShortMerc string)
	FzzFindMerchantByMerc(merchant *[]models.Merchant, Merc string)
	FzzFindMerchantByAddress(merchant *[]models.Merchant, Address string)
	FzzFindMerchantByNation(merchant *[]models.Merchant, Nation string)
	FzzFindMerchantByWebsite(merchant *[]models.Merchant, Website string)
	FzzFindMerchantByTaxType(merchant *[]models.Merchant, TaxType string)
	FzzFindMerchantByMercType(merchant *[]models.Merchant, MercType string)
	FzzFindMerchantByNotes(merchant *[]models.Merchant, Notes string)

	// Cust 相关接口
	FindCustByName(custs *[]models.Cust, name string)
	FindCustByGender(custs *[]models.Cust, gender string)
	FindCustByNation(custs *[]models.Cust, nation string)
	FindCustByAddr(custs *[]models.Cust, addr string)
	FindCustByEmail(custs *[]models.Cust, email string)
	FindCustByPhoneNum(custs *[]models.Cust, phoneNum string)
	FindCustByQQ(custs *[]models.Cust, qq string)
	FindCustByWechat(custs *[]models.Cust, wechat string)
	FindCustByMerc(custs *[]models.Cust, merc string)
	FindCustByPost(custs *[]models.Cust, post string)
	FindCustByNotes(custs *[]models.Cust, notes string)
	DeleteCust(cust *models.Cust) error
	SaveCust(cust *models.Cust) error
	FindCust(cust *[]models.Cust)
	FindCustByCustId(cust *[]models.Cust, CustId uint)
	FzzFindCustByName(cust *[]models.Cust, Name string)
	FzzFindCustByNation(cust *[]models.Cust, Nation string)
	FzzFindCustByAddr(cust *[]models.Cust, Addr string)
	FzzFindCustByMerc(cust *[]models.Cust, Merc string)
	FzzFindCustByPost(cust *[]models.Cust, Post string)
	FzzFindCustByNotes(cust *[]models.Cust, Notes string)

	// BankAccount 相关接口
	FindBankAccountByBankAccName(bankAccounts *[]models.BankAccount, bankAccName string)
	FindBankAccountByCompName(bankAccounts *[]models.BankAccount, compName string)
	FindBankAccountByAcctNum(bankAccounts *[]models.BankAccount, acctNum string)
	FindBankAccountByCurrency(bankAccounts *[]models.BankAccount, currency string)
	FindBankAccountByBankName(bankAccounts *[]models.BankAccount, bankName string)
	FindBankAccountByBankEng(bankAccounts *[]models.BankAccount, bankEng string)
	FindBankAccountByBankNum(bankAccounts *[]models.BankAccount, bankNum string)
	FindBankAccountBySwiftCode(bankAccounts *[]models.BankAccount, swiftCode string)
	FindBankAccountByBankAddr(bankAccounts *[]models.BankAccount, bankAddr string)
	FindBankAccountByNotes(bankAccounts *[]models.BankAccount, notes string)
	DeleteBankAccount(bankAccount *models.BankAccount) error
	SaveBankAccount(bankAccount *models.BankAccount) error
	FindBankAccount(bankAccount *[]models.BankAccount)
	FindBankAccountByBankAccId(bankAccount *[]models.BankAccount, BankAccId uint)
	FzzFindBankAccountByBankAccName(bankAccount *[]models.BankAccount, BankAccName string)
	FzzFindBankAccountByCompName(bankAccount *[]models.BankAccount, CompName string)
	FzzFindBankAccountByCurrency(bankAccount *[]models.BankAccount, Currency string)
	FzzFindBankAccountByBankName(bankAccount *[]models.BankAccount, BankName string)
	FzzFindBankAccountByBankEng(bankAccount *[]models.BankAccount, BankEng string)
	FzzFindBankAccountByBankAddr(bankAccount *[]models.BankAccount, BankAddr string)
	FzzFindBankAccountByNotes(bankAccount *[]models.BankAccount, Notes string)
	FindCustByMercId(custs *[]models.Cust, MercId uint)
	// Cust 相关接口
	FzzFindCustByEmail(custs *[]models.Cust, email string)
	FzzFindCustByPhoneNum(custs *[]models.Cust, phoneNum string)
	FzzFindCustByQQ(custs *[]models.Cust, qq string)
	FzzFindCustByWechat(custs *[]models.Cust, wechat string)

	// BankAccount 相关接口
	FindBankAccountByMercId(bankAccounts *[]models.BankAccount, MercId uint)
}

type service struct {
	db     *sql.DB
	gormDB *gorm.DB
}

var (
	dbname     = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: newLogger,
	})
	// AutoMigrate all models
	gormDB.AutoMigrate(
		&models.User{},
		&models.Acct{},
		&models.AcctBank{},
		&models.File{},
		&models.SuprType{},
		&models.PrdtType{},
		&models.FoodAddType{},
		&models.FeedAddType{},
		&models.UnitMeas{},
		&models.PackType{},
		&models.ConType{},
		&models.Currency{},
		&models.TradeTerm{},
		&models.Nation{},
		&models.Port{},
		&models.TaxType{},
		&models.BrandType{},
		&models.Degree{},
		&models.Dept{},
		&models.Post{},
		&models.QualStd{},
		&models.InvLoc{},
		&models.DocReq{},
		&models.PayMth{},
		&models.PayLimit{},
		&models.FinaDocStatus{},
		&models.FinaDocType{},
		&models.ExpType{},
		&models.Rates{},
		&models.BussOrderSta{},
		&models.MercType{},
		&models.Merchant{},
		&models.Cust{},
		&models.BankAccount{},
	)

	dbInstance = &service{
		db:     db,
		gormDB: gormDB,
	}

	return dbInstance
}

func (s *service) CreateFile(file *models.File) error {
	return s.gormDB.Create(file).Error
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}
