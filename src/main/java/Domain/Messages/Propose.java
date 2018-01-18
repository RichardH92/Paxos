package Domain.Messages;

public class Propose implements Comparable<Propose> {
    private int number;
    private int id;

    public Propose(int number, int id) {
        this.number = number;
        this.id = id;
    }

    public int getId() {
        return id;
    }

    public int getNumber() {
        return number;
    }

    @Override
    public int compareTo(Propose propose) {
        return -1;
    }
}
