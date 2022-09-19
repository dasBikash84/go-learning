import example.simple.SimpleOuterClass.Simple

fun main(args: Array<String>) {
    val simple =
        Simple.newBuilder()
            .setIsSimple(true)
            .setId(10)
            .setName("bikash")
            .addSampleList(2)
            .addAllSampleList(4..10)
            .build()

    println(simple)
}