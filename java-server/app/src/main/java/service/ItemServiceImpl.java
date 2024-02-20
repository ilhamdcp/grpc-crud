package service;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import grpc.service.item.Item;
import grpc.service.item.ItemRequest;
import grpc.service.item.ItemResponse;
import grpc.service.item.ItemServiceGrpc.ItemServiceImplBase;
import io.grpc.stub.StreamObserver;

public class ItemServiceImpl extends ItemServiceImplBase {
    private List<Item> items = new ArrayList<>();
    private List<String> itemNames = List.of(
        "Salmon Maki",
        "Salmon Sashimi",
        "Ikura Sushi",
        "Unagi Kabayaki",
        "Maguro Sashimi",
        "Toro Sashimi"
    );

    private List<String> descriptions = List.of(
        "the classic rolls of sushi rice stuffed with raw fresh salmon and spring onion and wrapped in nori (seaweed).",
        "salmon sashimi is simply made by cutting the flesh into slices and dipping them into some kind of dipping sauce.",
        "Ikura is the Japanese word for salmon roe (fish eggs), which can also be referred to as red caviar",
        "grilled eel drenched in thick and sweet soy sauce makes anyone's mouth water.",
        "Maguro is the Japanese word for tuna. Tuna is a saltwater fish that has a mild taste and a tender texture. It contains a little bit of oil, keeping the fish moist for sushi.",
        " Toro is a Japanese term for the fatty underbelly portion of a tuna. With its delightful cool, meaty texture combined with a tender, buttery taste, it creates the perfect combination to delight the palette."
    );


    public ItemServiceImpl() {
        super();
        initData();
    }

    private void initData() {

        for (int i = 0; i < itemNames.size(); i++) {
            Item currentItem = Item.newBuilder()
            .setId(i+1)
            .setItemName(itemNames.get(i))
            .setDescription(descriptions.get(i))
            .build();
            items.add(currentItem);
        }
    }

    @Override
    public void getItemById(ItemRequest request, StreamObserver<ItemResponse> responseObserver) {
        Item foundItem = items.stream().filter(i -> request.getId() == i.getId()).findAny().orElse(null);
        ItemResponse itemResponse = ItemResponse.newBuilder().setMessage(foundItem != null ? "success" : "failed").setItem(foundItem).build();
        responseObserver.onNext(itemResponse);
        responseObserver.onCompleted();
    }

    @Override
    public void getItemsByName(ItemRequest request, StreamObserver<Item> responseObserver) {
        List<Item> foundItems = items.stream().filter(i -> i.getItemName().toLowerCase().contains(request.getName())).collect(Collectors.toList());
        for (Item i : foundItems) {
            responseObserver.onNext(i);
        }
        responseObserver.onCompleted();
    }
}
