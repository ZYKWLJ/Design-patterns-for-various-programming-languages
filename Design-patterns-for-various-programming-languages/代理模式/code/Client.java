// import java.sql.Time;

public class Client {
    private Subject subject=new Agent(new Host());
    public static void main(String[] args) throws InterruptedException {
        Client client = new Client();
        // Thread.sleep(5000);
        client.subject.rent();   
    }
}
